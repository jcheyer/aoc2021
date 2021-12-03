package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay03(t *testing.T) {
	d3 := Day03{}
	err := d3.Load("testfiles/d3demo.txt")
	assert.NoError(t, err)
	assert.Equal(t, 5, d3.reportLen)
	assert.Equal(t, 12, d3.reportCount)
	assert.Equal(t, []int{7, 5, 8, 7, 5}, d3.on)
	s := d3.Part1()
	assert.Equal(t, "198", s)
	assert.Equal(t, "10110", d3.gamma)
	assert.Equal(t, "01001", d3.epsilon)

	most, least := d3.mostAndLeastCommonRuneAtPos(d3.rawInput, 0)
	assert.Equal(t, uint8('1'), most)
	assert.Equal(t, uint8('0'), least)

	most, least = d3.mostAndLeastCommonRuneAtPos(d3.rawInput, 1)
	assert.Equal(t, uint8('0'), most)
	assert.Equal(t, uint8('1'), least)

	// "fake"-Test
	s = d3.Part2()
	assert.Equal(t, "230", s)
}
