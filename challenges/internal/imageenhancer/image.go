package imageenhancer

import (
	"strconv"
)

type coord struct {
	x, y int
}

type image struct {
	pixels     map[coord]bool
	maxX, maxY int
	minX, minY int
	generation int
	bitchMode  bool
}

func New(data []string) *image {
	l := len(data)
	image := &image{
		pixels: make(map[coord]bool),
		maxX:   l,
		maxY:   l,
		minX:   0,
		minY:   0,
	}

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {
			image.pixels[coord{x, y}] = data[y][x] == '#'
		}
	}

	return image
}

func (i *image) IsOn(x, y int) bool {
	isOn, ok := i.pixels[coord{x, y}]
	if ok {
		return isOn
	}
	if i.bitchMode {
		return i.generation%2 == 1
	}

	return isOn
}

func (i *image) IsOnStr(x, y int) string {
	if i.IsOn(x, y) {
		return "1"
	}
	return "0"
}

func (i *image) BitchMode() {
	i.bitchMode = true
}

func (i *image) Enhance(algo string) *image {
	if algo[0] == '#' {
		i.bitchMode = true
	}
	output := &image{
		pixels:     make(map[coord]bool),
		maxX:       i.maxX + 1,
		maxY:       i.maxY + 1,
		minX:       i.minX - 1,
		minY:       i.minY - 1,
		generation: i.generation + 1,
		bitchMode:  i.bitchMode,
	}

	for x := i.minX - 1; x <= i.maxX+1; x++ {
		for y := i.minY - 1; y <= i.maxY+1; y++ {
			surround := i.PixelSurrounding(x, y)
			output.pixels[coord{x, y}] = algo[surround] == '#'
		}
	}

	return output
}

func (i *image) PixelSurrounding(x, y int) int {
	bin := ""
	bin += i.IsOnStr(x-1, y-1)
	bin += i.IsOnStr(x, y-1)
	bin += i.IsOnStr(x+1, y-1)
	bin += i.IsOnStr(x-1, y)
	bin += i.IsOnStr(x, y)
	bin += i.IsOnStr(x+1, y)
	bin += i.IsOnStr(x-1, y+1)
	bin += i.IsOnStr(x, y+1)
	bin += i.IsOnStr(x+1, y+1)
	ret, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(ret)
}

func (i *image) Count() int {
	c := 0
	for _, p := range i.pixels {
		if p {
			c++
		}
	}
	return c
}
