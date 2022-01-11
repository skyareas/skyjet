package main

import (
	"github.com/akaahmedkamal/go-cli/cmd"
	"github.com/akaahmedkamal/go-server/app"
	"github.com/akaahmedkamal/go-server/cmd/db"
	"github.com/akaahmedkamal/go-server/cmd/server"
)

func main() {
	// get app instance
	a := app.Shared()

	// register db commands
	a.Register(&db.Init{})
	a.Register(&db.Migrate{})

	// register server commands
	a.Register(&server.Start{})

	// register default commands
	a.Register(&cmd.Help{})
	a.Register(&cmd.Version{})

	// start the app
	a.Run()
}
