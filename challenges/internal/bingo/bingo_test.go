package bingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()
	assert.False(t, c.loaded)
	assert.False(t, c.HasBingo())
}
