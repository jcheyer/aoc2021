package graph

import (
	"testing"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/stretchr/testify/assert"
)

func TestCaveWorld(t *testing.T) {
	t1, err := lib.ReadInputLines("testfiles/t1.txt")
	assert.NoError(t, err)
	cw := New(t1)
	assert.Len(t, cw.nodes, 6)
	assert.Len(t, cw.nodes["start"], 2)
	assert.Len(t, cw.nodes["b"], 4)
	assert.Equal(t, 10, cw.CountPossiblePaths())
}
