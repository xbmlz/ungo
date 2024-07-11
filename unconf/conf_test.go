package unconf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AllConf struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"database"`
}

func TestNewWithPath(t *testing.T) {
	var c AllConf

	config, err := New("./testdata/config.yaml")
	assert.NoError(t, err)
	assert.NotNil(t, config)

	err = config.Parse(&c)
	assert.NoError(t, err)

	server := config.Get("server")
	assert.NotNil(t, server)

	assert.Equal(t, c.Server.Port, 8080)
	assert.Equal(t, c.Database.Host, "localhost")
}
