package snailfishmath

import (
	"testing"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	data, err := lib.ReadInputLines("./../../testfiles/d18demo.txt")
	assert.NoError(t, err)

	n := New(data[0])
	for i := 1; i < len(data); i++ {
		add := New(data[i])
		n = n.Add(add)
	}
	expected := "[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]"
	assert.Equal(t, expected, n.String())
	assert.Equal(t, 4140, n.Magnitude())
}
