package ungo

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/xbmlz/ungo/log"
	"github.com/xbmlz/ungo/server"
)

// App is the main application struct.
type App struct {
	id      string
	name    string
	version string
	servers []server.Server
	signals []os.Signal
}

// Option is a function that can be passed to NewApp to modify the App's behavior.
type Option func(application *App)

func NewApp(opts ...Option) *App {
	app := &App{}
	for _, op := range opts {
		op(app)
	}
	// default random id
	if len(app.id) == 0 {
		bytes := make([]byte, 24)
		_, _ = rand.Read(bytes)
		app.id = hex.EncodeToString(bytes)
	}
	// default accept signals
	if len(app.signals) == 0 {
		app.signals = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	}
	return app
}

// WithID application add id
func WithID(id string) func(application *App) {
	return func(application *App) {
		application.id = id
	}
}

// WithName application add name
func WithName(name string) func(application *App) {
	return func(application *App) {
		application.name = name
	}
}

// WithVersion application add version
func WithVersion(version string) func(application *App) {
	return func(application *App) {
		application.version = version
	}
}

// WithServer application add server
func WithServer(servers ...server.Server) func(application *App) {
	return func(application *App) {
		application.servers = servers
	}
}

// WithSignals application add listen signals
func WithSignals(signals []os.Signal) func(application *App) {
	return func(application *App) {
		application.signals = signals
	}
}

// Run application run
func (app *App) Run(ctx context.Context) error {
	if len(app.servers) == 0 {
		return nil
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, app.signals...)
	errCh := make(chan error, 1)

	for _, s := range app.servers {
		go func(srv server.Server) {
			if err := srv.Start(); err != nil {
				log.Errorf("failed to start server, err: %s", err)
				errCh <- err
			}
		}(s)
	}

	select {
	case err := <-errCh:
		_ = app.Stop()
		return err
	case <-ctx.Done():
		return app.Stop()
	case <-quit:
		return app.Stop()
	}
}

// Stop application stop
func (app *App) Stop() error {
	wg := sync.WaitGroup{}
	for _, s := range app.servers {
		wg.Add(1)
		go func(srv server.Server) {
			defer wg.Done()
			if err := srv.Shutdown(); err != nil {
				log.Errorf("failed to stop server, err: %s", err)
			}
		}(s)
	}
	// wait all server graceful shutdown
	wg.Wait()
	return nil
}
