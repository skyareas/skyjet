package db

import "github.com/akaahmedkamal/go-cli/v1"

// Init command to initialize the database.
type Init struct {
	Name string `cli:"name"`
	Help string `cli:"help"`
}

// NewInitCmd initialized a new db/init command.
func NewInitCmd() *Init {
	return &Init{
		Name: "db/init",
		Help: "initialize database",
	}
}

// Run executes the command's logic.
func (i *Init) Run(app *cli.App) {
	app.Log().Error().Fatal("not implemented!")
}
