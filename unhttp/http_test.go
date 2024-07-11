package unhttp

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// new a http http.Handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	server := New(handler)

	go func() {
		time.Sleep(2 * time.Second)

		resp, err := http.Get("http://localhost:8080")
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, http.StatusOK)

		time.Sleep(1 * time.Second)
		err = server.ShutdownWithTimeout(2 * time.Second)
		assert.NoError(t, err)
		t.Log("shutdown completed")
	}()

	_ = server.Start()
}
