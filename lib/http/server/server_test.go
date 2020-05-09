package server

import (
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("custom-server", func(t *testing.T) {
		_ = ListenAndServe(":8080", nil)
	})
}