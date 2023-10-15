package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

const ConfigFile = "ja.toml"
const VendorDir = "vendor"
const OutDir = "out"
const SrcDir = "src"
const MainClass = "app.App"

type Config struct {
	Packages []string
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
