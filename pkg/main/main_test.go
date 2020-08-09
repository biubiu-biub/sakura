package main

import (
	"github.com/biubiu-biub/sakura/pkg/entry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntry(t *testing.T) {
	assert.Equal(t, entry.Ping(), "ping", "The two words should be the same.")
}
