package db

import (
	"github.com/akaahmedkamal/go-cli/v1"
)

type Init struct{}

func (i *Init) Name() string {
	return "db/init"
}

func (i *Init) Desc() string {
	return "initialize database"
}

func (i *Init) Run(app *cli.App) {
}
