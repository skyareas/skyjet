package config

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/akaahmedkamal/go-server/app"
	"gopkg.in/yaml.v2"
)

// Config struct represents the app's configurations.
type Config struct {
	Http HttpConfig `yaml:"Http"`
	Db   DbConfig   `yaml:"DB"`
}

// configFilePath returns full path to the config file.
// first, it looks at the cmd-line args, if not found,
// it returns the default config file path.
func configFilePath() string {
	p, exists := app.Shared().Args().LookupString("--cfg", "--config")
	if exists {
		return resolve(p)
	}
	return resolve(DefaultConfigFilePath)
}

// readConfigFile reads and parses the content of config file,
// and returns the result as a generic map type, see CfgFile.
func readConfigFile() (*Config, error) {
	cfg := Config{}
	d, err := os.ReadFile(configFilePath())
	if err != nil {
		return &cfg, err
	}
	err = yaml.Unmarshal(d, &cfg)
	return &cfg, err
}

// mustConvToInt is a helper function to convert string to int,
// it will panic when fails to convert.
func mustConvToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

// dbDriver returns db driver name, it looks up
// the driver name in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value.
func dbDriver(cfg *Config) string {
	args := app.Shared().Args()

	driver, found := args.LookupString("--db-driver")
	if found {
		return driver
	}

	if cfg.Db.Driver != "" {
		return cfg.Db.Driver
	}

	return DefaultDbDriver
}

// dbUrl returns db url, it looks up
// the url in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value.
func dbUrl(cfg *Config) string {
	args := app.Shared().Args()

	url, found := args.LookupString("--db-url")
	if found {
		return url
	}

	if cfg.Db.Url != "" {
		return cfg.Db.Url
	}

	return DefaultDbUrl
}

// httpHost returns http host name, it looks up
// the host name in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value.
func httpHost(cfg *Config) string {
	args := app.Shared().Args()

	host, found := args.LookupString("--http-host")
	if found {
		return host
	}

	if cfg.Http.Host != "" {
		return cfg.Http.Host
	}

	return DefaultHttpHost
}

// httpPort returns http port, it looks up
// the port in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value.
func httpPort(cfg *Config) int {
	args := app.Shared().Args()

	port, found := args.LookupString("--http-port")
	if found {
		return mustConvToInt(port)
	}

	if cfg.Http.Port != 0 {
		return cfg.Http.Port
	}

	return DefaultHttpPort
}

// httpReadTimeout returns http read timeout, it looks up
// the read-timeout in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value, see DefaultHttpReadTimeout.
func httpReadTimeout(cfg *Config) time.Duration {
	args := app.Shared().Args()

	t, found := args.LookupString("--http-read-time")
	if found {
		return time.Duration(mustConvToInt(t)) * time.Second
	}

	if cfg.Http.ReadTimeout != 0 {
		return cfg.Http.ReadTimeout * time.Second
	}

	return DefaultHttpReadTimeout
}

// httpWriteTimeout returns http write timeout, it looks up
// the write-timeout in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value, see DefaultHttpWriteTimeout.
func httpWriteTimeout(cfg *Config) time.Duration {
	args := app.Shared().Args()

	t, found := args.LookupString("--http-write-time")
	if found {
		return time.Duration(mustConvToInt(t)) * time.Second
	}

	if cfg.Http.WriteTimeout != 0 {
		return cfg.Http.WriteTimeout * time.Second
	}

	return DefaultHttpWriteTimeout
}

// httpIdleTimeout returns http idle timeout, it looks up
// the idle-timeout in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value, see DefaultHttpIdleTimeout.
func httpIdleTimeout(cfg *Config) time.Duration {
	args := app.Shared().Args()

	t, found := args.LookupString("--http-idle-time")
	if found {
		return time.Duration(mustConvToInt(t)) * time.Second
	}

	if cfg.Http.IdleTimeout != 0 {
		return cfg.Http.IdleTimeout * time.Second
	}

	return DefaultHttpIdleTimeout
}

// httpViewsPath returns http views root path, it looks up
// the views-path in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value, see DefaultHttpViewsPath.
func httpViewsPath(cfg *Config) string {
	args := app.Shared().Args()

	vp, found := args.LookupString("--http-views-path")
	if found {
		return resolve(vp)
	}

	if cfg.Http.ViewsPath != "" {
		return resolve(cfg.Http.ViewsPath)
	}

	return resolve(DefaultHttpViewsPath)
}

// httpSessionName returns http session name, it looks up
// the session name in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value.
func httpSessionName(cfg *Config) string {
	args := app.Shared().Args()

	name, found := args.LookupString("--http-session-name")
	if found {
		return name
	}

	if cfg.Http.Session.CookieName != "" {
		return cfg.Http.Session.CookieName
	}

	return DefaultHttpSessionName
}

// httpSessionSecret returns http session secret, it looks up
// the session secret in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value.
func httpSessionSecret(cfg *Config) string {
	args := app.Shared().Args()

	secret, found := args.LookupString("--http-session-secret")
	if found {
		return secret
	}

	if cfg.Http.Session.Secret != "" {
		return cfg.Http.Session.Secret
	}

	return DefaultHttpSessionSecret
}

// resolve returns an absolute path from given path.
func resolve(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	wd, _ := os.Getwd()
	return filepath.Join(wd, path)
}
