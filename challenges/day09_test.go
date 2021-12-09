package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay09(t *testing.T) {
	d9 := Day09{}
	err := d9.Load("testfiles/d09demo.txt")
	assert.NoError(t, err)
	assert.Len(t, d9.field, 7)
	assert.Len(t, d9.field[0], 12)
	assert.Equal(t, "15", d9.Part1())
	assert.Equal(t, "1134", d9.Part2())
	/*assert.Equal(t, 3, d9.acquireBasin(1, 1))
	assert.Equal(t, 0, d9.acquireBasin(1, 1))
	assert.Equal(t, 9, d9.acquireBasin(10, 1))
	assert.Equal(t, 14, d9.acquireBasin(1, 4))
	assert.Equal(t, 0, d9.acquireBasin(1, 4))*/

}
