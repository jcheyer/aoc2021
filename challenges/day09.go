package challenges

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/jcheyer/aoc2021/lib"
)

type Day09 struct {
	rawInput []string
	field    [][]uint8
	basin    [][]uint8
}

func (d *Day09) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	d.field, err = parseLines2Heights(d.rawInput)
	if err != nil {
		return err
	}

	d.basin, err = parseLines2Heights(d.rawInput)
	if err != nil {
		return err
	}

	return nil
}

func (d *Day09) Part1() string {
	risk := 0
	//fmt.Printf("fdsafdsa")

	//fmt.Printf("Todo:\nX: %d Y: %d\n", len(d.field[0])-2, len(d.field)-2)
	for y := 1; y < len(d.field)-1; y++ {
		for x := 1; x < len(d.field[0])-1; x++ {
			val := d.field[y][x]
			//fmt.Printf("X: %d Y: %d Value: %d\n", x, y, val)

			if d.field[y][x+1] > val &&
				d.field[y-1][x] > val &&
				d.field[y+1][x] > val &&
				d.field[y][x-1] > val {
				risk = risk + int(val) + 1
				//fmt.Printf("Hit.... X: %d Y: %d Value: %d Risk: %d\n", x, y, val, risk)
			}

		}
	}

	return fmt.Sprintf("%d", risk)
}

func (d *Day09) Part2() string {
	var basins []int
	for y := 1; y < len(d.field)-1; y++ {
		for x := 1; x < len(d.field[0])-1; x++ {
			if d.basin[y][x] < 9 {
				size := d.acquireBasin(x, y)
				if size > 0 {
					basins = append(basins, size)
				}
			}
		}
	}
	sort.Ints(basins)
	count := len(basins)
	top3 := basins[count-1] * basins[count-2] * basins[count-3]

	return fmt.Sprintf("%d", top3)
}

func (d *Day09) acquireBasin(x, y int) int {
	size := 0
	if d.basin[y][x] > 8 {
		return size
	}
	d.basin[y][x] = 99
	size = size + 1
	size = size + d.acquireBasin(x, y+1)
	size = size + d.acquireBasin(x-1, y)
	size = size + d.acquireBasin(x+1, y)
	size = size + d.acquireBasin(x, y-1)

	return size
}

func parseLines2Heights(lines []string) ([][]uint8, error) {
	width := len(lines[0])
	field := make([][]uint8, 0)

	border := make([]uint8, 0)
	for i := 0; i < width+2; i++ {
		border = append(border, uint8(math.MaxUint8))
	}
	field = append(field, border)

	for _, line := range lines {
		row := make([]uint8, width+1)
		row[0] = uint8(math.MaxUint8)
		for i, r := range line {
			x, err := strconv.Atoi(string(r))
			if err != nil {
				return field, err
			}

			row[i+1] = uint8(x)
			if err != nil {
				return field, err
			}
		}
		row = append(row, uint8(math.MaxUint8))
		field = append(field, row)
	}

	field = append(field, border)

	return field, nil
}
