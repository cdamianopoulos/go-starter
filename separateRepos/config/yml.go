// Package config implements utility functions for loading configuration data from various sources.
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// LoadYaml expects parameter obj to be a pointer.
// E.g: err := config.LoadYaml("config.yml", &env)
func LoadYaml(fileName string, obj interface{}) error {
	src, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(src, obj)
}

// MustLoadYaml expects parameter obj to be a pointer.
// E.g: config.MustLoadYaml("config.yml", &env)
func MustLoadYaml(fileName string, obj interface{}) {
	if err := LoadYaml(fileName, obj); err != nil {
		panic(err)
	}
}
