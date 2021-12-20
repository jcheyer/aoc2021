package scanner

import (
	"testing"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	data, err := lib.ReadInputLines("./../../testfiles/d19demo.txt")
	assert.NoError(t, err)

	sc := New(data)
	assert.Len(t, sc, 5)

}
