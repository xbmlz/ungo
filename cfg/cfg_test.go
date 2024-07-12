package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	var c struct {
		Server struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"server"`
	}

	err := Load("./testdata/config.yaml", &c)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	assert.Equal(t, c.Server.Port, 8080)
	assert.Equal(t, c.Server.Host, "0.0.0.0")
}
