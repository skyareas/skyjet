package main

import (
	"os"

	cliCmd "github.com/akaahmedkamal/go-cli/cmd"
	"github.com/akaahmedkamal/go-cli/v1"
	dbCmd "github.com/akaahmedkamal/go-server/cmd/db"
	srvCmd "github.com/akaahmedkamal/go-server/cmd/server"
	"github.com/akaahmedkamal/go-server/config"
	"github.com/akaahmedkamal/go-server/db"
)

func main() {
	// create app instance
	app := cli.NewApp(os.Args[1:])

	// create config instance for this app
	cfg := config.New(app)

	// set db vars
	dbInstance, err := db.Connect(
		cfg.DbDriver(),
		cfg.DbUrl(),
	)
	if err != nil {
		panic(err)
	}
	app.Set("db", dbInstance)

	// set http vars
	app.Set("http/host", cfg.HttpHost())
	app.Set("http/port", cfg.HttpPort())

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
