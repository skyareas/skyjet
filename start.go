package skyjet

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/skyareas/go-cli"
)

// Start command to start the Http server.
type Start struct {
	Name       string `cli:"name"`
	Help       string `cli:"help"`
	ConfigFile string `cli:"option" optName:"config" optHelp:"Path to config file"`
	stop       chan bool
	quit       chan os.Signal
}

// NewStartCmd initialized a new server/start command.
func NewStartCmd() *Start {
	return &Start{
		Name: "start",
		Help: "Starts the Http server",
		stop: make(chan bool, 1),
		quit: make(chan os.Signal, 1),
	}
}

// Run executes the command's logic.
func (s *Start) Run(_ *cli.App) {
	if s.ConfigFile != "" {
		err := app.LoadConfigFile(s.ConfigFile)
		if err != nil {
			app.log.Fatalln(err)
		}
	}
	signal.Notify(s.quit, os.Interrupt)

	go s.shutdown()
	go app.srv.ListenAndServe()

	<-s.stop

	if db != nil {
		_ = db.Disconnect()
	}
}

// shutdown gracefully shuts down the Http server.
func (s *Start) shutdown() {
	<-s.quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.srv.Shutdown(ctx); err != nil {
		panic(err.Error())
	}

	close(s.stop)
}
