package skyjet

import (
	"github.com/sirupsen/logrus"
	"github.com/skyareas/go-cli"
)

type Application struct {
	cli.App
	Router
	cfg *Config
	srv *HttpServer
	log *logrus.Logger
}

var app *Application

func App() *Application {
	if app == nil {
		app = &Application{*cli.NewApp(), *NewRouter(), nil, nil, NewJsonLogger()}
		app.cfg, _ = loadConfigFile(defaultConfigFilePath, true)
		app.srv = NewHttpServer(&app.Router)
		app.Register(NewStartCmd())
		app.SetDefaultCommand("start")
	}
	return app
}

func (a *Application) Config() *Config {
	return a.cfg
}

func (a *Application) LoadConfigFile(path string) error {
	var err error
	a.cfg, err = loadConfigFile(path)
	return err
}

func (a *Application) Log() *logrus.Logger {
	return a.log
}
