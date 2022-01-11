package db

import "github.com/akaahmedkamal/go-cli/v1"

// Init command to initialize the database.
type Init struct{}

// Name returns the command name.
func (i *Init) Name() string {
	return "db/init"
}

// Desc returns the command description.
func (i *Init) Desc() string {
	return "initialize database"
}

// Run executes the command's logic.
func (i *Init) Run(app *cli.App) {
	app.Log().Error().Fatal("not implemented!")
}
