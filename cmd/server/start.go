package server

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/akaahmedkamal/go-cli/v1"
	"github.com/akaahmedkamal/go-server/app"
	"github.com/akaahmedkamal/go-server/config"
	"github.com/akaahmedkamal/go-server/db"
	"github.com/akaahmedkamal/go-server/server"
	"github.com/akaahmedkamal/go-server/server/middleware"
	"github.com/akaahmedkamal/go-server/server/routes"
	"github.com/akaahmedkamal/go-server/server/routes/auth"
	"github.com/akaahmedkamal/go-server/server/routes/home"
	"github.com/akaahmedkamal/go-server/server/routes/users"
)

// Start command to start the Http server.
type Start struct {
	Name string `cli:"name"`
	Help string `cli:"help"`
	srv  *server.HttpServer
	stop chan bool
	quit chan os.Signal
}

// NewStartCmd initialized a new server/start command.
func NewStartCmd() *Start {
	return &Start{
		Name: "server/start",
		Help: "start Http server",
	}
}

// Run executes the command's logic.
func (s *Start) Run(app *cli.App) {
	// make sure to close the db
	// connection before exising
	defer closeDb()

	s.stop = make(chan bool, 1)
	s.quit = make(chan os.Signal, 1)
	s.srv = setupHttpServer()

	signal.Notify(s.quit, os.Interrupt)

	go s.shutdown()
	go s.srv.ListenAndServe()

	cfg := config.Shared()
	app.Log().Infof("[HTTP]: server started at %s:%d\n", cfg.Http.Host, cfg.Http.Port)

	<-s.stop

	app.Log().Infoln("[HTTP]: server stopped.")
}

// shutdown gracefully shuts down the Http server.
func (s *Start) shutdown() {
	<-s.quit

	app.Shared().Log().Infoln("[HTTP]: shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		app.Shared().Log().Fatalf("[HTTP]: %s\n", err.Error())
	}

	close(s.stop)
}

// setupHttpServer initializes and configures
// a new http server instance.
func setupHttpServer() *server.HttpServer {
	// create http server instance,
	// with default router
	srv := server.NewHttpServer()

	// get ref to the router
	r := srv.Router()

	// register middleware
	r.Use("*", &middleware.Cors{
		Origin:  []string{"*"},
		Methods: []string{"*"},
		Headers: []string{"*"},
	})

	// register home routes
	r.Get("/", &home.Index{})
	r.Get("/about", &home.About{})

	// register auth routes
	r.All("/auth/register", &auth.Register{})
	r.All("/auth/login", &auth.Login{})
	r.Post("/auth/logout", &auth.Logout{})

	// register user routes
	r.Get("/users/:id/profile", &users.Profile{})

	// register not found route
	r.All("*", &routes.NotFound{})

	return srv
}

// closeDb closes the database connection if found.
func closeDb() {
	if err := db.Disconnect(); err != nil {
		app.Shared().Log().Errorf("[DB]: %s\n", err.Error())
	}
}
