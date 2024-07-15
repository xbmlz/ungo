package server

import "strconv"

// Server is transport server.
type Server interface {
	Start() error
	Shutdown() error
}

type Config struct {
	Address string `json:"address" yaml:"address" env:"SERVER_ADDRESS" default:"0.0.0.0"`
	Port    int    `json:"port" yaml:"port" env:"SERVER_PORT" default:"8080"`
}

func (c *Config) Addr() string {
	return c.Address + ":" + strconv.Itoa(c.Port)
}
