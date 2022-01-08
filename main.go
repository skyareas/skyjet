package main

import (
	"log"
	"os"

	cliCmd "github.com/akaahmedkamal/go-cli/cmd"
	"github.com/akaahmedkamal/go-cli/v1"
	dbCmd "github.com/akaahmedkamal/go-server/cmd/db"
	srvCmd "github.com/akaahmedkamal/go-server/cmd/server"
	"github.com/akaahmedkamal/go-server/config"
	"github.com/akaahmedkamal/go-server/db"
	"github.com/akaahmedkamal/go-server/server"
	"github.com/akaahmedkamal/go-server/server/routes/auth"
)

func main() {
	// create app instance
	app := cli.NewApp(os.Args[1:])

	// setup app modules
	app.Set("cfg", config.New(app))
	app.Set("db", db.New(app))
	app.Set("http", setupHttpServer(app))

	// make sure to close the db
	// connection before exising
	defer closeDb()

	// register db commands
	app.Register(&dbCmd.Init{})
	app.Register(&dbCmd.Migrate{})

	// register server commands
	app.Register(&srvCmd.Start{})

	// register default commands
	app.Register(&cliCmd.Help{})
	app.Register(&cliCmd.Version{})

	// start the app
	app.Run()
}

func setupHttpServer(app *cli.App) *server.HttpServer {
	// create http server instance
	srv := server.NewHttpServer(app)

	// get ref to the router
	r := srv.Router()

	// register auth routes
	r.Get("/login", &auth.Login{})
	r.Get("/logout", &auth.Logout{})
	r.Get("/register", &auth.Register{})

	return srv
}

func closeDb() {
	if err := db.Disconnect(); err != nil {
		log.Fatalf("[DB]: %s\n", err.Error())
	}
}
