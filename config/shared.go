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
		shared = new(Config)

		file, err := readConfigFile()
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			app.Shared().Log().Fatalf("failed to read config file: %s", err.Error())
		}

		shared.Http = &HttpConfig{
			Host:         httpHost(file),
			Port:         httpPort(file),
			ReadTimeout:  httpReadTimeout(file),
			WriteTimeout: httpWriteTimeout(file),
			IdleTimeout:  httpIdleTimeout(file),
			ViewsPath:    httpViewsPath(file),
		}

		shared.Db = &DbConfig{
			Driver: dbDriver(file),
			Url:    dbUrl(file),
		}
	}

	return shared
}
