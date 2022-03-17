package skyjet

import (
	"github.com/sirupsen/logrus"
	"github.com/skyareas/go-cli"
)

type App struct {
	cli.App
	Router
	cfg *Config
	srv *HttpServer
	log *logrus.Logger
}

var app *App

func SharedApp() *App {
	if app == nil {
		app = &App{*cli.NewApp(), *NewRouter(), nil, nil, logrus.New()}
		app.cfg, _ = loadConfigFile(defaultConfigFilePath, true)
		app.srv = NewHttpServer(&app.Router)
		app.Register(NewStartCmd())
		app.SetDefaultCommand("start")
	}
	return app
}

func (a *App) Config() *Config {
	return a.cfg
}

func (a *App) LoadConfigFile(path string) error {
	var err error
	a.cfg, err = loadConfigFile(path)
	return err
}
