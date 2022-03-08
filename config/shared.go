package config

import (
	"errors"
	"os"

	"github.com/akaahmedkamal/go-server/app"
)

// shared instance of the app config.
var shared *Config

// Shared returns the shared/global instance of the
// app config if found, otherwise, it initializes
// a new shared instance and returns it.
func Shared() *Config {
	if shared == nil {
		shared = &Config{}

		cfg, err := readConfigFile()
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			app.Shared().Log().Fatalf("failed to read config file: %s", err.Error())
		}

		shared.Http = HttpConfig{
			Host:         httpHost(cfg),
			Port:         httpPort(cfg),
			ReadTimeout:  httpReadTimeout(cfg),
			WriteTimeout: httpWriteTimeout(cfg),
			IdleTimeout:  httpIdleTimeout(cfg),
			ViewsPath:    httpViewsPath(cfg),
			Session: HttpSessionConfig{
				CookieName: httpSessionName(cfg),
				Secret:     httpSessionSecret(cfg),
			},
		}

		shared.Db = DbConfig{
			Driver: dbDriver(cfg),
			Url:    dbUrl(cfg),
		}
	}

	return shared
}
