package config

// DbConfig struct represents the database config values.
type DbConfig struct {
	Driver string `yaml:"Driver"`
	Url    string `yaml:"Url"`
}
