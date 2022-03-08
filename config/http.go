package config

import "time"

// HttpConfig struct represents the http-server config values.
type HttpConfig struct {
	Host         string            `yaml:"Host"`
	Port         int               `yaml:"Port"`
	ViewsPath    string            `yaml:"ViewsPath"`
	ReadTimeout  time.Duration     `yaml:"ReadTimeout"`
	WriteTimeout time.Duration     `yaml:"WriteTimeout"`
	IdleTimeout  time.Duration     `yaml:"IdleTimeout"`
	Session      HttpSessionConfig `yaml:"Session"`
}

type HttpSessionConfig struct {
	CookieName string `yaml:"CookieName"`
	Secret     string `yaml:"Secret"`
}
