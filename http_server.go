package skyjet

import (
	"context"
	"fmt"
	"net/http"
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
		srv.router = NewRouter()
	}

	addr := fmt.Sprintf("%s:%d", app.cfg.Http.Host, app.cfg.Http.Port)

	srv.srv = &http.Server{
		Addr:         addr,
		Handler:      srv.router,
		ReadTimeout:  app.cfg.Http.ReadTimeout,
		WriteTimeout: app.cfg.Http.WriteTimeout,
		IdleTimeout:  app.cfg.Http.IdleTimeout,
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
	app.log.Printf("starting http server at %s", s.srv.Addr)
	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err.Error())
	}
}

// Shutdown stop listening to new incoming requests,
// and finalize all live requests if any.
func (s *HttpServer) Shutdown(ctx context.Context) error {
	s.srv.SetKeepAlivesEnabled(false)
	err := s.srv.Shutdown(ctx)
	if err != nil {
		app.log.Printf("http server stopped with error: %s", err.Error())
	} else {
		app.log.Println("http server stopped successfully")
	}
	return err
}
