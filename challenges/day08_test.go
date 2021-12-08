package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay08(t *testing.T) {
	d8 := Day08{}
	err := d8.Load("testfiles/d08demo.txt")
	assert.NoError(t, err)
	assert.Equal(t, "26", d8.Part1())

	assert.Equal(t, "61229", d8.Part2())
}
