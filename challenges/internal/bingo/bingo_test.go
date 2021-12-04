package bingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []string{
	"22 13 17 11  0",
	"8  2 23  4 24",
	"21  9 14 16  7",
	"6 10  3 18  5",
	"1 12 20 15 19",
}

func TestNew(t *testing.T) {
	c := New()
	assert.False(t, c.loaded)
	assert.False(t, c.HasBingo())
	err := c.Setup(testData)
	assert.NoError(t, err)
	// Stichprobe
	assert.Equal(t, int64(15), c.Fields[4][3])
}
