package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/akaahmedkamal/go-cli/v1"
	"github.com/akaahmedkamal/go-server/config"
)

type HttpServer struct {
	app    *cli.App
	srv    *http.Server
	router *Router
}

func NewHttpServer(app *cli.App) *HttpServer {
	cfg := config.Of(app)

	srv := new(HttpServer)

	srv.app = app
	srv.srv = new(http.Server)
	srv.router = NewRouter(app)

	srv.srv.Addr = fmt.Sprintf("%s:%d", cfg.HttpHost(), cfg.HttpPort())
	srv.srv.Handler = srv.router
	srv.srv.ReadTimeout = cfg.HttpReadTimeout()
	srv.srv.WriteTimeout = cfg.HttpWriteTimeout()
	srv.srv.IdleTimeout = cfg.HttpIdleTimeout()

	return srv
}

func (s *HttpServer) Router() *Router {
	return s.router
}

func (s *HttpServer) ListenAndServe() error {
	err := s.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	s.srv.SetKeepAlivesEnabled(false)
	return s.srv.Shutdown(ctx)
}
