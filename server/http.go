package server

import (
	"context"
	"errors"
	"net/http"
	"time"
)

var _ Server = (*HTTPServer)(nil)

type HTTPServer struct {
	srv *http.Server
}

func NewHTTPServer(handler http.Handler, config *Config) *HTTPServer {
	ser := HTTPServer{
		srv: &http.Server{
			Addr:    config.Addr(),
			Handler: handler,
		},
	}

	return &ser
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
