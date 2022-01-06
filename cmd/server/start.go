package server

import (
	"log"

	"github.com/akaahmedkamal/go-cli/v1"
	"github.com/akaahmedkamal/go-server/server"
)

type Start struct {
	srv *server.HttpServer
}

func (s *Start) Name() string {
	return "server/start"
}

func (s *Start) Desc() string {
	return "start server"
}

func (s *Start) Run(app *cli.App) {
	s.srv = server.NewHttpServer()

	defer s.Shutdown()

	if err := s.srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func (s *Start) Shutdown() {
	if err := s.srv.Shutdown(); err != nil {
		log.Fatal(err)
	}
}
