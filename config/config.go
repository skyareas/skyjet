package config

import (
	"github.com/akaahmedkamal/go-cli/v1"
)

const DefaultHttpPort = "5000"
const DefaultDbDriver = "sqite3"
const DefaultDbUrl = "db.sql"

type Config struct {
	app *cli.App
}

func New(app *cli.App) *Config {
	return &Config{app}
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
	host, _ := c.app.Args().GetString("-h", "--host")
	return host
}

func (c *Config) HttpPort() string {
	port, _ := c.app.Args().GetString("-p", "--port")
	if port == "" {
		port = DefaultHttpPort
	}
	return port
}
