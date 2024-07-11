package unhttp

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Port            int           `json:"port" yaml:"port" env:"HTTP_PORT" default:"8080"`
	Host            string        `json:"host" yaml:"host" env:"HTTP_HOST" default:"0.0.0.0"`
	ShutdownTimeout time.Duration `json:"shutdown_timeout" yaml:"shutdown_timeout" env:"HTTP_SHUTDOWN_TIMEOUT" default:"10"`
}

type Server struct {
	srv *http.Server
}

type Options func(*Config)

var config = Config{
	Port:            8080,
	Host:            "0.0.0.0",
	ShutdownTimeout: 10,
}

func New(handler http.Handler, options ...Options) *Server {

	for _, option := range options {
		option(&config)
	}

	server := &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
			Handler: handler,
		},
	}
	return server
}

func (s *Server) Start() (err error) {
	err = s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Shutdown() (err error) {
	err = s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func WithPort(port int) Options {
	return func(c *Config) {
		c.Port = port
	}
}

func WithHost(host string) Options {
	return func(c *Config) {
		c.Host = host
	}
}

func WithShutdownTimeout(timeout time.Duration) Options {
	return func(c *Config) {
		c.ShutdownTimeout = timeout
	}
}
