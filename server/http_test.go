package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPServer(t *testing.T) {
	srv := NewHTTPServer(Config{
		Port: 8080,
		Host: "localhost",
	})
	assert.NotNil(t, srv)
}
