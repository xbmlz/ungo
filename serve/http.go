package serve

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Port int    `json:"port" yaml:"port" env:"HTTP_PORT" default:"8080"`
	Host string `json:"host" yaml:"host" env:"HTTP_HOST" default:"0.0.0.0"`
}

type HTTPServer struct {
	Srv    *http.Server
	Router *gin.Engine
}

func NewHTTPServer(config Config) (srv *HTTPServer, err error) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/ping", func(ctx *gin.Context) { ctx.String(200, "OK") })

	srv = &HTTPServer{
		Srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
			Handler: r,
		},
		Router: r,
	}

	return srv, nil
}

func MustNewHTTPServer(config Config) *HTTPServer {
	srv, err := NewHTTPServer(config)
	if err != nil {
		log.Fatalf("Failed to create HTTP server: %v", err)
	}
	return srv
}

func (s *HTTPServer) Run() {
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := s.Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
