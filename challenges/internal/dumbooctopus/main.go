package dumbooctopus

import "strconv"

type dumbooctopus struct {
	field        [][]byte
	sizeX, sizeY int
}

func New(data []string) *dumbooctopus {

	d := &dumbooctopus{
		sizeX: len(data[0]),
		sizeY: len(data),
	}

	field := make([][]byte, 0)
	for i := 0; i < d.sizeY; i++ {
		row := []byte{}
		for _, r := range data[i] {
			octo, _ := strconv.Atoi(string(r))
			row = append(row, byte(octo))
		}
		field = append(field, row)

	}

	d.field = field

	return d
}

func (d *dumbooctopus) Step() int {

	for y := 0; y < d.sizeY; y++ {
		for x := 0; x < d.sizeX; x++ {
			d.field[y][x]++
		}
	}

	count := 0
	for y := 0; y < d.sizeY; y++ {
		for x := 0; x < d.sizeX; x++ {
			count += d.Flash(x, y)
		}
	}

	return count
}

func (d *dumbooctopus) Flash(x, y int) int {
	if d.field[y][x] < 10 {
		return 0
	}

	d.field[y][x] = 0
	count := 1
	for yy := y - 1; yy <= y+1; yy++ {
		for xx := x - 1; xx <= x+1; xx++ {
			if xx < 0 || yy < 0 || // out of bounds
				yy >= d.sizeY || xx >= d.sizeX || // out of bounds
				d.field[yy][xx] == 0 { // already flashed
				continue
			}
			d.field[yy][xx]++
			count += d.Flash(xx, yy)
		}
	}
	return count
}

func (d *dumbooctopus) Sum() int {
	sum := 0
	for y := 0; y < d.sizeY; y++ {
		for x := 0; x < d.sizeX; x++ {
			sum += int(d.field[y][x])
		}
	}
	return sum
}
