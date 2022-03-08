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
	Http *HttpConfig `yaml:"Http"`
	Db   *DbConfig   `yaml:"DB" json:"DB"`
}

// CfgFile generic type represents the content of a config file.
type CfgFile map[interface{}]interface{}

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
func readConfigFile() (CfgFile, error) {
	d, err := os.ReadFile(configFilePath())
	if err != nil {
		return nil, err
	}
	raw := make(CfgFile)
	err = yaml.Unmarshal(d, raw)
	return raw, err
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
func dbDriver(cfgFile CfgFile) string {
	args := app.Shared().Args()

	driver, found := args.LookupString("--db-driver")
	if found {
		return driver
	}

	if db, ok := cfgFile["DB"].(CfgFile); ok {
		if driver, ok = db["Driver"].(string); ok {
			return driver
		}
	}

	return DefaultDbDriver
}

// dbUrl returns db url, it looks up
// the url in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value.
func dbUrl(cfgFile CfgFile) string {
	args := app.Shared().Args()

	url, found := args.LookupString("--db-url")
	if found {
		return url
	}

	if db, ok := cfgFile["DB"].(CfgFile); ok {
		if url, ok = db["Url"].(string); ok {
			return url
		}
	}

	return DefaultDbUrl
}

// httpHost returns http host name, it looks up
// the host name in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value.
func httpHost(cfgFile CfgFile) string {
	args := app.Shared().Args()

	host, found := args.LookupString("--http-host")
	if found {
		return host
	}

	if http, ok := cfgFile["Http"].(CfgFile); ok {
		if host, ok = http["Host"].(string); ok {
			return host
		}
	}

	return DefaultHttpHost
}

// httpPort returns http port, it looks up
// the port in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value.
func httpPort(cfgFile CfgFile) int {
	args := app.Shared().Args()

	port, found := args.LookupString("--http-port")
	if found {
		return mustConvToInt(port)
	}

	if http, ok := cfgFile["Http"].(CfgFile); ok {
		if port, ok := http["Port"].(int); ok {
			return port
		}
	}

	return DefaultHttpPort
}

// httpReadTimeout returns http read timeout, it looks up
// the read-timeout in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value, see DefaultHttpReadTimeout.
func httpReadTimeout(cfgFile CfgFile) time.Duration {
	args := app.Shared().Args()

	t, found := args.LookupString("--http-read-time")
	if found {
		return time.Duration(mustConvToInt(t)) * time.Second
	}

	if http, ok := cfgFile["Http"].(CfgFile); ok {
		if t, ok := http["ReadTimeout"].(int); ok {
			return time.Duration(t) * time.Second
		}
	}

	return DefaultHttpReadTimeout
}

// httpWriteTimeout returns http write timeout, it looks up
// the write-timeout in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value, see DefaultHttpWriteTimeout.
func httpWriteTimeout(cfgFile CfgFile) time.Duration {
	args := app.Shared().Args()

	t, found := args.LookupString("--http-write-time")
	if found {
		return time.Duration(mustConvToInt(t)) * time.Second
	}

	if http, ok := cfgFile["Http"].(CfgFile); ok {
		if t, ok := http["WriteTimeout"].(int); ok {
			return time.Duration(t) * time.Second
		}
	}

	return DefaultHttpWriteTimeout
}

// httpIdleTimeout returns http idle timeout, it looks up
// the idle-timeout in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value, see DefaultHttpIdleTimeout.
func httpIdleTimeout(cfgFile CfgFile) time.Duration {
	args := app.Shared().Args()

	t, found := args.LookupString("--http-idle-time")
	if found {
		return time.Duration(mustConvToInt(t)) * time.Second
	}

	if http, ok := cfgFile["Http"].(CfgFile); ok {
		if t, ok := http["IdleTimeout"].(int); ok {
			return time.Duration(t) * time.Second
		}
	}

	return DefaultHttpIdleTimeout
}

// httpViewsPath returns http views root path, it looks up
// the views-path in the following order.
// - 1st looks at the cmd-line args.
// - 2nd looks at the config file.
// - 3rd returns the default hard-coded value, see DefaultHttpViewsPath.
func httpViewsPath(cfgFile CfgFile) string {
	args := app.Shared().Args()

	vp, found := args.LookupString("--http-views-path")
	if found {
		return resolve(vp)
	}

	if http, ok := cfgFile["Http"].(CfgFile); ok {
		if vp, ok = http["ViewsPath"].(string); ok {
			return resolve(vp)
		}
	}

	return resolve(DefaultHttpViewsPath)
}

// resolve returns an absolute path from given path.
func resolve(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	wd, _ := os.Getwd()
	return filepath.Join(wd, path)
}
