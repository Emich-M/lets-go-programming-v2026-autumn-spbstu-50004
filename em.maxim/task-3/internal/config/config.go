package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Settings, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Settings{}, err
	}

	var cfg Settings
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return Settings{}, err
	}

	return cfg, nil
}
