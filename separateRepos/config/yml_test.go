package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadYaml(t *testing.T) {
	const yamlContents = "port: 14556\nnewrelic: abc456"

	f, err := os.CreateTemp("", "")
	assert.NoError(t, err)

	defer os.Remove(f.Name())

	l, err := f.WriteString(yamlContents)
	assert.NoError(t, err, "failed to write string to temporary file")
	assert.Equal(t, len(yamlContents), l)

	assert.NoError(t, f.Close(), "failed to close file")

	config := struct {
		Port     uint16
		NewRelic string
	}{}

	err = LoadYaml(f.Name(), &config)
	assert.NoError(t, err, "failed to load YAML config file")

	assert.Equal(t, uint16(14556), config.Port)
	assert.Equal(t, "abc456", config.NewRelic)
}

func TestLoadYamlEmpty(t *testing.T) {
	const yamlContents = "::::::::::::::::"

	f, err := os.CreateTemp("", "")
	assert.NoError(t, err)

	defer os.Remove(f.Name())

	l, err := f.WriteString(yamlContents)
	assert.NoError(t, err, "failed to write string to temporary file")
	assert.Equal(t, len(yamlContents), l)

	assert.NoError(t, f.Close(), "failed to close file")

	config := struct {
		Port     uint16
		NewRelic string
	}{}

	err = LoadYaml(f.Name(), &config)
	assert.NoError(t, err, "failed to load YAML config file")

	assert.Empty(t, config.Port)
	assert.Empty(t, config.NewRelic)
}

func TestLoadYamlFails(t *testing.T) {
	config := struct {
		Port     uint16
		NewRelic string
	}{}

	err := LoadYaml("z_y_x_w_v", &config)
	assert.Error(t, err, "expected LoadYaml() to return an error")
	assert.Empty(t, config.Port)
	assert.Empty(t, config.NewRelic)
}

func TestMustLoadYaml(t *testing.T) {
	const yamlContents = "port: 14556\nnewrelic: abc456"

	f, err := os.CreateTemp("", "")
	assert.NoError(t, err)

	defer os.Remove(f.Name())

	l, err := f.WriteString(yamlContents)
	assert.NoError(t, err, "failed to write string to temporary file")
	assert.Equal(t, len(yamlContents), l)

	assert.NoError(t, f.Close(), "failed to close file")

	config := struct {
		Port     uint16
		NewRelic string
	}{}

	MustLoadYaml(f.Name(), &config)
	assert.Equal(t, uint16(14556), config.Port)
	assert.Equal(t, "abc456", config.NewRelic)
}

func TestMustLoadYamlPanics(t *testing.T) {
	assert.Panics(t,
		func() {
			m := map[string]string{}
			MustLoadYaml("z_y_x_w_v", &m)
		},
		"expected MustLoadYaml to panic when provided with invalid YAML",
	)
}
