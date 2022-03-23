package skyjet

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"time"
)

const (
	defaultConfigFilePath    = "config.json"
	defaultHttpHost          = "0.0.0.0"
	defaultHttpPort          = 8080
	defaultHttpReadTimeout   = 5 * time.Second
	defaultHttpWriteTimeout  = 10 * time.Second
	defaultHttpIdleTimeout   = 15 * time.Second
	defaultHttpViewsPath     = "views"
	defaultHttpContentRoot   = "public"
	defaultHttpSessionName   = "session"
	defaultHttpSessionSecret = "abc@123#xyz"
	defaultDbDriver          = "sqlite3"
	defaultDbUrl             = "db.sql"
)

// Config struct represents the app's configurations.
type Config struct {
	Http         HttpConfig
	Db           DbConfig
	CustomConfig map[string]interface{}
}

// HttpConfig struct represents the http-server config values.
type HttpConfig struct {
	Host         string
	Port         int
	ViewsPath    string
	ContentRoot  string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	Session      HttpSessionConfig
}

// HttpSessionConfig struct represents the http-session config values.
type HttpSessionConfig struct {
	CookieName string
	Secret     string
}

// DbConfig struct represents the database config values.
type DbConfig struct {
	Driver string
	Url    string
}

// loadConfigFile reads and parses the content of config file.
func loadConfigFile(path string, ignoringErrors ...bool) (*Config, error) {
	cfg := Config{}

	var ignoreErrors bool
	if len(ignoringErrors) != 0 {
		ignoreErrors = ignoringErrors[0]
	}

	d, err := ioutil.ReadFile(path)
	if err != nil && !ignoreErrors {
		return &cfg, err
	}

	ext := filepath.Ext(path)
	switch ext {
	case ".json":
		err = json.Unmarshal(d, &cfg)
	}

	cfg.applyDefaults()

	return &cfg, err
}

// applyDefaults applies default values to the empty config values.
func (c *Config) applyDefaults() {
	if c.Http.Host == "" {
		c.Http.Host = defaultHttpHost
	}
	if c.Http.Port == 0 {
		c.Http.Port = defaultHttpPort
	}
	if c.Http.ReadTimeout == 0 {
		c.Http.ReadTimeout = defaultHttpReadTimeout
	}
	if c.Http.WriteTimeout == 0 {
		c.Http.WriteTimeout = defaultHttpWriteTimeout
	}
	if c.Http.IdleTimeout == 0 {
		c.Http.IdleTimeout = defaultHttpIdleTimeout
	}
	if c.Http.ViewsPath == "" {
		c.Http.ViewsPath = defaultHttpViewsPath
	}
	if c.Http.ContentRoot == "" {
		c.Http.ContentRoot = defaultHttpContentRoot
	}
	if c.Http.Session.CookieName == "" {
		c.Http.Session.CookieName = defaultHttpSessionName
	}
	if c.Http.Session.Secret == "" {
		c.Http.Session.Secret = defaultHttpSessionSecret
	}
	if c.Db.Driver == "" {
		c.Db.Driver = defaultDbDriver
	}
	if c.Db.Url == "" {
		c.Db.Url = defaultDbUrl
	}
}
