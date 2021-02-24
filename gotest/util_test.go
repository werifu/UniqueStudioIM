package gotest

import (
	"crypto/rand"
	"io"
	"testing"
)

func Test_sessionId(t *testing.T) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return
	}
	t.Log(b)
}
