package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/akaahmedkamal/go-server/app"
	"github.com/akaahmedkamal/go-server/config"
)

// HttpServer struct represents the
// Http server implementation.
type HttpServer struct {
	srv    *http.Server
	router *Router
}

// NewHttpServer initializes a new HttpServer and return it.
// optionally, pass a pointer to the Router instance to be
// used; if not passed a default router will be used.
func NewHttpServer(router ...*Router) *HttpServer {
	srv := new(HttpServer)

	if len(router) > 0 {
		srv.router = router[0]
	} else {
		srv.router = NewRouter(&RouterConfig{})
	}

	cfg := config.Shared()
	srv.srv = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port),
		Handler:      srv.router,
		ReadTimeout:  cfg.Http.ReadTimeout,
		WriteTimeout: cfg.Http.WriteTimeout,
		IdleTimeout:  cfg.Http.IdleTimeout,
	}

	return srv
}

// Router returns a pointer to the Http Router.
func (s *HttpServer) Router() *Router {
	return s.router
}

// ListenAndServe start listening at the address specified,
// and handles incoming requests.
func (s *HttpServer) ListenAndServe() {
	err := s.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		app.Shared().Log().Fatalf("[HTTP]: %s\n", err.Error())
	}
}

// Shutdown stop listening to new incoming requests,
// and finalize all live requests if any.
func (s *HttpServer) Shutdown(ctx context.Context) error {
	s.srv.SetKeepAlivesEnabled(false)
	return s.srv.Shutdown(ctx)
}
