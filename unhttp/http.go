package unhttp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Port int    `json:"port" yaml:"port" env:"HTTP_PORT" default:"8080"`
	Host string `json:"host" yaml:"host" env:"HTTP_HOST" default:"0.0.0.0"`
}

type Server struct {
	srv *http.Server
}

type Options func(*Config)

var config = Config{
	Port: 8080,
	Host: "0.0.0.0",
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
	return s.srv.Shutdown(context.Background())
}

func (s *Server) ShutdownWithTimeout(timeout time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return s.srv.Shutdown(ctx)
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