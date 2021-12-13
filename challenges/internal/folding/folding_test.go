package folding

import (
	"testing"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/stretchr/testify/assert"
)

func TestFolding(t *testing.T) {
	data, err := lib.ReadInputLines("./../../testfiles/d13demo.txt")
	assert.NoError(t, err)
	paper := New(data)
	assert.Len(t, paper.dots, 18)
	assert.True(t, paper.dots[coord{x: 6, y: 10}])
	assert.False(t, paper.dots[coord{x: 5432, y: 1430}])
	//assert.Equal(t, 10, paper.dots[0].y)

	assert.Len(t, paper.Folds, 2)
	paper.Fold(paper.Folds[0])

	assert.Len(t, paper.dots, 17)

}
