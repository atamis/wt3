package config

import (
	"github.com/BurntSushi/toml"
)

type configuration struct {
	TemplatePath string
}

var Config configuration

func init() {

	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		// handle error
		panic(err)
	}

}
