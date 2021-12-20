package imageenhancer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []string{
	"#..#.",
	"#....",
	"##..#",
	"..#..",
	"..###",
}

func TestNewImage(t *testing.T) {

	i := New(data)
	//assert.Len(t, i.pixels, 10)
	assert.Equal(t, 10, i.Count())
	isOn, ok := i.pixels[coord{0, 0}]
	assert.True(t, isOn)
	assert.True(t, ok)

	isOn, ok = i.pixels[coord{0, 4}]
	assert.False(t, isOn)
	assert.True(t, ok)

	assert.True(t, i.IsOn(0, 0))
	assert.False(t, i.IsOn(0, 4))
	assert.False(t, i.IsOn(100, 100))
	assert.False(t, i.IsOn(-100, -100))

}

func TestSurrounding(t *testing.T) {
	i := New(data)
	val := i.PixelSurrounding(2, 2)
	assert.Equal(t, 34, val)
}

func TestPart1(t *testing.T) {
	algo := "..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#"
	i := New(data)
	i = i.Enhance(algo)
	//assert.Len(t, i.pixels, 24)
	assert.Equal(t, 24, i.Count())
	i = i.Enhance(algo)
	//assert.Len(t, i.pixels, 35)
	assert.Equal(t, 35, i.Count())
}

func TestBitchMode(t *testing.T) {
	i := New(data)
	val := i.PixelSurrounding(0, 0)
	assert.Equal(t, 18, val)
	assert.Equal(t, "0", i.IsOnStr(-1, -1))
	i.BitchMode()

	val = i.PixelSurrounding(0, 0)
	assert.Equal(t, 18, val)
	i.generation = 1
	val = i.PixelSurrounding(0, 0)
	assert.Equal(t, 502, val)

}
