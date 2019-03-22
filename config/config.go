package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var appenv string
var cfgfile string

// Config type for app
type Config struct {
	Web      Web    `toml:"web"`
	GitHub   GitHub `toml:"github"`
	Database Database
}

func init() {
	env := configFile()

	// set cfgfile name
	cfgfile = fmt.Sprintf("config.%s.toml", env)
}

// NewConfig create an instance of app config.
// It returns a Config and an error
func NewConfig() (Config, error) {
	var cfg Config

	if _, err := toml.DecodeFile(cfgfile, &cfg); err != nil {
		log.Fatal(err)
		return cfg, errors.New("cannot open and decode file correctly")
	}

	return cfg, nil
}

// configFile returns name of config file based on environment variable.
// It returns current config file
func configFile() string {
	// set working environment
	appenv = os.Getenv("COMMIT_ENV")
	if (appenv != "development") &&
		(appenv != "test") &&
		(appenv != "production") {
		appenv = "development"
	}

	return appenv
}
