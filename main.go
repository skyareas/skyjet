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
	a.Register(db.NewInitCmd())
	a.Register(db.NewMigrateCmd())

	// register server commands
	a.Register(server.NewStartCmd())

	// register default commands
	a.Register(cmd.NewHelpCmd())
	a.Register(cmd.NewVersionCmd())

	// start the app
	a.Run()
}
