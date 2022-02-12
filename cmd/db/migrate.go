package db

import "github.com/akaahmedkamal/go-cli/v1"

// Migrate command to migrate the database.
type Migrate struct {
	Name string `cli:"name"`
	Help string `cli:"help"`
}

// NewMigrateCmd initialized a new db/migrate command.
func NewMigrateCmd() *Migrate {
	return &Migrate{
		Name: "db/migrate",
		Help: "migrate database",
	}
}

// Run executes the command's logic.
func (m *Migrate) Run(app *cli.App) {
	app.Log().Error().Fatal("not implemented!")
}
