package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var _ Server = (*HTTPServer)(nil)

type Config struct {
	Port int    `json:"port" yaml:"port" env:"HTTP_PORT" default:"8080"`
	Host string `json:"host" yaml:"host" env:"HTTP_HOST" default:"0.0.0.0"`
}

type HTTPServer struct {
	srv    *http.Server
	Router *gin.Engine
}

func NewHTTPServer(config Config) *HTTPServer {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/ping", func(ctx *gin.Context) { ctx.String(200, "OK") })

	srv := &HTTPServer{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
			Handler: r,
		},
		Router: r,
	}

	return srv
}

// Start to start the server and wait for it to listen on the given address
func (s *HTTPServer) Start() (err error) {
	err = s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Shutdown shuts down the server and close with graceful shutdown duration
func (s *HTTPServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
