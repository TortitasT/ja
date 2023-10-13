package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

const ConfigFile = "ja.toml"

type Config struct {
	Libraries []string
}

func LoadConfig() (Config, error) {
	var config Config

	filebytes, err := os.ReadFile(ConfigFile)
	if err != nil {
		return config, err
	}

	err = toml.Unmarshal(filebytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
