package vents

import (
	"testing"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/stretchr/testify/assert"
)

func TestField(t *testing.T) {
	f := New(10, 10)

	testData := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
	}
	f.AddLines(testData, false)
	t.Logf("%+v\n", f.Vents)

	assert.Equal(t, int64(0), f.Vents[0][0])
	assert.Equal(t, int64(1), f.Vents[0][9])
	assert.Equal(t, 0, f.VentCountGreaterOne(1))
	f.Vents[7][7] = 5
	f.Vents[8][8] = 5
	assert.Equal(t, 2, f.VentCountGreaterOne(1))
	t.Logf("%+v\n", f.Vents)
}

func TestExtractCoordinates(t *testing.T) {
	x1, y1, x2, y2, dir := extractCoordinates("0,9 -> 5,9")
	assert.Equal(t, 0, x1)
	assert.Equal(t, 9, y1)
	assert.Equal(t, 5, x2)
	assert.Equal(t, 9, y2)
	assert.Equal(t, "horizontal", dir)

	x1, y1, x2, y2, dir = extractCoordinates("8,0 -> 0,8")
	assert.Equal(t, 8, x1)
	assert.Equal(t, 0, y1)
	assert.Equal(t, 0, x2)
	assert.Equal(t, 8, y2)
	assert.Equal(t, "diagonal", dir)

	x1, y1, x2, y2, dir = extractCoordinates("7,0 -> 7,4")
	assert.Equal(t, 7, x1)
	assert.Equal(t, 0, y1)
	assert.Equal(t, 7, x2)
	assert.Equal(t, 4, y2)
	assert.Equal(t, "vertical", dir)
}

func TestLowerFirst(t *testing.T) {
	a, b := lowerFirst(5, 7)
	assert.Equal(t, 5, a)
	assert.Equal(t, 7, b)

	a, b = lowerFirst(7, 5)
	assert.Equal(t, 5, a)
	assert.Equal(t, 7, b)
}

func TestDemoP1(t *testing.T) {
	f := New(10, 10)
	data, err := lib.ReadInputLines("./../../testfiles/d05demo.txt")
	assert.NoError(t, err)
	f.AddLines(data, false)

	assert.Equal(t, 5, f.VentCountGreaterOne(1))
	assert.Equal(t, int64(2), f.Vents[0][9])
	assert.Equal(t, int64(2), f.Vents[1][9])

}

func TestDemoP2(t *testing.T) {
	f := New(10, 10)
	data, err := lib.ReadInputLines("./../../testfiles/d05demo.txt")
	assert.NoError(t, err)
	f.AddLines(data, true)
	t.Logf("%s", f.String())
	assert.Equal(t, 12, f.VentCountGreaterOne(1))
	assert.Equal(t, int64(1), f.Vents[0][0])
	assert.Equal(t, int64(1), f.Vents[2][0])

}
