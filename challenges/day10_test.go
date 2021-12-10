package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay10(t *testing.T) {
	d := Day10{}
	err := d.Load("testfiles/d10demo.txt")
	assert.NoError(t, err)
	assert.Len(t, d.rawInput, 10)
	assert.Equal(t, "26397", d.Part1())
	assert.Equal(t, "288957", d.Part2())
}
