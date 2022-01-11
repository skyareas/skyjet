package config

import "time"

const (
	DefaultConfigFilePath = "cfg.yml"
)

// Http Defaults
const (
	DefaultHttpHost         = "127.0.0.1"
	DefaultHttpPort         = 5000
	DefaultHttpReadTimeout  = 5 * time.Second
	DefaultHttpWriteTimeout = 10 * time.Second
	DefaultHttpIdleTimeout  = 15 * time.Second
	DefaultHttpViewsPath    = "views"
)

// Db Defaults
const (
	DefaultDbDriver = "sqlite3"
	DefaultDbUrl    = "db.sql"
)
