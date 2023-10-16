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
	Packages []string `toml:"packages"`
}

func NewConfig() (Config, error) {
	config := Config{
		Packages: []string{},
	}

	file, err := os.Create(ConfigFile)
	if err != nil {
		return config, err
	}

	defer file.Close()

	err = toml.NewEncoder(file).Encode(config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func ConfigFileExists() bool {
	_, err := os.Stat(ConfigFile)
	return !os.IsNotExist(err)
}

func LoadConfig() (Config, error) {
	var config Config

	if !ConfigFileExists() {
		return config, os.ErrNotExist
	}

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
