package unhttp

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xbmlz/ungo/unlog"
)

type Config struct {
	Port int    `json:"port" yaml:"port" env:"HTTP_PORT" default:"8080"`
	Host string `json:"host" yaml:"host" env:"HTTP_HOST" default:"0.0.0.0"`
}

type Server struct {
	srv *http.Server
}

func NewServer(handler http.Handler, config Config) *Server {
	server := &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
			Handler: handler,
		},
	}
	return server
}

func (s *Server) Run() (err error) {
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
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

	unlog.Infof("Shutting down server...")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		return errors.New("Server forced to shutdown: " + err.Error())
	}

	unlog.Infof("Server exiting")
	return nil
}
