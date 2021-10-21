package config

import "github.com/kelseyhightower/envconfig"

// MustLoadYamlEnv attempts to load YAML configuration file in yamlFilePath and then environment variables into obj.
// If yamlFilePath is empty, then no YAML file is parsed.
// obj is expected to be a pointer.
func MustLoadYamlEnv(yamlFilePath, envPrefix string, obj interface{}) {
	if yamlFilePath != "" {
		// Load defined fields from a YAML config file.
		MustLoadYaml(yamlFilePath, obj)
	}

	// Any environment variables loaded are over-riding any values defined from the YAML config file.
	envconfig.MustProcess(envPrefix, obj)
}

// MustLoadEnvYaml attempts to load environment variables and then YAML configuration file in yamlFilePath into obj. MustLoadEnvYaml is the same as MustLoadYamlEnv except it is executed in reverse order.
// If yamlFilePath is empty, then no YAML file is parsed.
// obj is expected to be a pointer.
func MustLoadEnvYaml(envPrefix, yamlFilePath string, obj interface{}) {
	// Load available environment variables into obj.
	envconfig.MustProcess(envPrefix, obj)

	// Any fields defined in the YAML config file override environment variables.
	if yamlFilePath != "" {
		// Load defined fields from a YAML config file.
		MustLoadYaml(yamlFilePath, obj)
	}
}
