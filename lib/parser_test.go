package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInputLines(t *testing.T) {
	expected := []string{"157", "148", "149", "146", "144", "145", "162", "163", "164", "166"}
	s1, err := ReadInputLines("testfiles/input.txt")
	assert.NoError(t, err)
	assert.Len(t, s1, 10)
	assert.Equal(t, expected, s1)

	_, err = ReadInputLines("not found.txt")
	assert.Error(t, err)
}

func TestParseLines2Int64(t *testing.T) {
	input := []string{"157", "148", "149", "146", "144", "145", "162", "163", "164", "166"}
	expected := []int64{157, 148, 149, 146, 144, 145, 162, 163, 164, 166}
	output, err := ParseLines2Int(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, output)

	_, err = ParseLines2Int([]string{"not a number", "123"})
	assert.Error(t, err)
}
