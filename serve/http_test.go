package serve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPServer(t *testing.T) {
	_, err := NewHTTPServer(Config{
		Port: 8080,
		Host: "localhost",
	})
	assert.Nil(t, err)
}
