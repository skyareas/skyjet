package config

import "time"

const (
	DefaultConfigFilePath = "config.yml"
)

// Http Defaults
const (
	DefaultHttpHost          = "0.0.0.0"
	DefaultHttpPort          = 5000
	DefaultHttpReadTimeout   = 5 * time.Second
	DefaultHttpWriteTimeout  = 10 * time.Second
	DefaultHttpIdleTimeout   = 15 * time.Second
	DefaultHttpViewsPath     = "views"
	DefaultHttpSessionName   = "session"
	DefaultHttpSessionSecret = "abc@123#xyz"
)

// Db Defaults
const (
	DefaultDbDriver = "sqlite3"
	DefaultDbUrl    = "db.sql"
)
