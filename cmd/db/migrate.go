package db

import "github.com/akaahmedkamal/go-cli/v1"

// Migrate command to migrate the database.
type Migrate struct{}

// Name returns the command name.
func (m *Migrate) Name() string {
	return "db/migrate"
}

// Desc returns the command description.
func (m *Migrate) Desc() string {
	return "migrate database"
}

// Run executes the command's logic.
func (m *Migrate) Run(app *cli.App) {
	app.Log().Error().Fatal("not implemented!")
}
