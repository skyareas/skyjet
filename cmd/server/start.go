package server

import (
	"context"
	"github.com/akaahmedkamal/go-server/config"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/akaahmedkamal/go-cli/v1"
	"github.com/akaahmedkamal/go-server/server"
)

type Start struct {
	srv  *server.HttpServer
	done chan bool
	quit chan os.Signal
}

func (s *Start) Name() string {
	return "server/start"
}

func (s *Start) Desc() string {
	return "start server"
}

func (s *Start) Run(app *cli.App) {
	s.done = make(chan bool, 1)
	s.quit = make(chan os.Signal, 1)
	s.srv = app.Get("http").(*server.HttpServer)

	signal.Notify(s.quit, os.Interrupt)

	go s.shutdown()

	cfg := config.Of(app)
	log.Printf("[HTTP]: server started at %s:%d\n", cfg.HttpHost(), cfg.HttpPort())
	if err := s.srv.ListenAndServe(); err != nil {
		log.Fatalf("[HTTP]: %s\n", err.Error())
	}

	<-s.done

	log.Println("[HTTP]: server stopped.")
}

func (s *Start) shutdown() {
	<-s.quit
	log.Println("[HTTP]: shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatalf("[HTTP]: %s\n", err.Error())
	}

	close(s.done)
}
