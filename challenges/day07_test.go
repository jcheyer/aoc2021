package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD7(t *testing.T) {

	d7 := Day07{}
	err := d7.Load("testfiles/d07demo.txt")
	assert.NoError(t, err)
	result := d7.Part1()
	assert.Equal(t, "37", result)
	assert.Equal(t, uint64(41), d7.calcFuelP1(1))
	assert.Equal(t, uint64(39), d7.calcFuelP1(3))
	assert.Equal(t, uint64(71), d7.calcFuelP1(10))

	assert.Equal(t, uint64(206), d7.calcFuelP2(2))
	assert.Equal(t, "168", d7.Part2())
}
