package diracdice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImplements(t *testing.T) {
	dd := NewDDice(0)
	assert.Implements(t, (*dice)(nil), dd)
}

func TestDDice(t *testing.T) {
	ddice := NewDDice(0)
	assert.Equal(t, 1, ddice.Roll())
	assert.Equal(t, 2, ddice.Roll())
	assert.Equal(t, 3, ddice.Roll())
	assert.Equal(t, 4, ddice.Roll())
	ddice.lastDraw = 98
	assert.Equal(t, 99, ddice.Roll())
	assert.Equal(t, 100, ddice.Roll())
	assert.Equal(t, 1, ddice.Roll())
	assert.Equal(t, 2, ddice.Roll())
	assert.Equal(t, 8, ddice.UseCount())
}
