package config

import (
	"strconv"
	"time"

	"github.com/akaahmedkamal/go-cli/v1"
)

// Http Defaults
const (
	DefaultHttpPort         = 5000
	DefaultHttpReadTimeout  = 5 * time.Second
	DefaultHttpWriteTimeout = 10 * time.Second
	DefaultHttpIdleTimeout  = 15 * time.Second
)

// Db Defaults
const (
	DefaultDbDriver = "sqlite3"
	DefaultDbUrl    = "db.sql"
)

type Config struct {
	app *cli.App
}

func New(app *cli.App) *Config {
	return &Config{app}
}

func Of(app *cli.App) *Config {
	cfg, ok := app.Get("cfg").(*Config)
	if !ok || cfg == nil {
		panic("unable to find module \"cfg\"")
	}
	return cfg
}

func (c *Config) DbDriver() string {
	dbDriver, _ := c.app.Args().GetString("--db-driver")
	if dbDriver == "" {
		dbDriver = DefaultDbDriver
	}
	return dbDriver
}

func (c *Config) DbUrl() string {
	dbUrl, _ := c.app.Args().GetString("--db-url")
	if dbUrl == "" {
		dbUrl = DefaultDbUrl
	}
	return dbUrl
}

func (c *Config) HttpHost() string {
	host, _ := c.app.Args().GetString("--http-host")
	return host
}

func (c *Config) HttpPort() int {
	port, _ := c.app.Args().GetString("--http-port")
	if port == "" {
		return DefaultHttpPort
	}
	return c.MustConvToInt(port)
}

func (c *Config) HttpReadTimeout() time.Duration {
	t, _ := c.app.Args().GetString("--http-read-time")
	if t == "" {
		return DefaultHttpReadTimeout
	}
	return time.Duration(c.MustConvToInt(t)) * time.Second
}

func (c *Config) HttpWriteTimeout() time.Duration {
	t, _ := c.app.Args().GetString("--http-write-time")
	if t == "" {
		return DefaultHttpWriteTimeout
	}
	return time.Duration(c.MustConvToInt(t)) * time.Second
}

func (c *Config) HttpIdleTimeout() time.Duration {
	t, _ := c.app.Args().GetString("--http-idle-time")
	if t == "" {
		return DefaultHttpIdleTimeout
	}
	return time.Duration(c.MustConvToInt(t)) * time.Second
}

func (c *Config) MustConvToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
