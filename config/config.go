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
	env := ConfigFile()

	// set cfgfile name
	cfgfile = fmt.Sprintf("config.%s.toml", env)
}

func NewConfig() (Config, error) {
	var cfg Config

	if _, err := toml.DecodeFile(cfgfile, &cfg); err != nil {
		log.Fatal(err)
		return cfg, errors.New("cannot open and decode file correctly.")
	}

	return cfg, nil
}

func ConfigFile() string {
	// set working environment
	appenv = os.Getenv("COMMIT_ENV")
	if (appenv != "development") &&
		(appenv != "test") &&
		(appenv != "production") {
		appenv = "development"
	}

	return appenv
}
