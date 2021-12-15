package pathfinder

import (
	"testing"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/stretchr/testify/assert"
)

func TestDay15(t *testing.T) {
	data, err := lib.ReadInputLines("./../../testfiles/d15demo.txt")
	assert.NoError(t, err)
	pf, err := New(data, 5)
	assert.NoError(t, err)
	assert.NotNil(t, pf)
	assert.Equal(t, 40, pf.Best1())
	assert.Equal(t, 315, pf.Best2())
	//spew.Dump(pf.d)
	assert.Equal(t, uint8(1), val(pf.data, 0, 0))
	assert.Equal(t, uint8(2), val(pf.data, 10, 0))
	assert.Equal(t, uint8(2), val(pf.data, 9, 0))
	assert.Equal(t, uint8(3), val(pf.data, 19, 0))
	assert.Equal(t, uint8(6), val(pf.data, 49, 0))
	assert.Equal(t, uint8(9), val(pf.data, 49, 49))
	assert.Equal(t, uint8(1), val(pf.data, 9, 9))

}

func TestCoord2Vertex(t *testing.T) {
	data, err := lib.ReadInputLines("./../../testfiles/d15demo.txt")
	assert.NoError(t, err)
	pf, err := New(data, 5)
	assert.NoError(t, err)
	assert.Equal(t, 0, pf.coord2vertex(0, 0))
	assert.Equal(t, 49, pf.coord2vertex(49, 0))
	assert.Equal(t, 50, pf.coord2vertex(0, 1))
	assert.Equal(t, 51, pf.coord2vertex(1, 1))

	assert.Equal(t, uint8(7), val(pf.data, 0, 4))

	assert.Equal(t, 200, pf.coord2vertex(0, 4))
	assert.Equal(t, 306, pf.coord2vertex(6, 6))
}
