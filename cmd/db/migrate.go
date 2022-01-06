package db

import (
	"github.com/akaahmedkamal/go-cli/v1"
)

type Migrate struct{}

func (m *Migrate) Name() string {
	return "db/migrate"
}

func (m *Migrate) Desc() string {
	return "migrate database"
}

func (m *Migrate) Run(app *cli.App) {
}
