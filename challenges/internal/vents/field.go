package vents

import (
	"fmt"
	"strconv"
	"strings"
)

type Field struct {
	Vents      [][]int64
	readyToUse bool
}

func New(x, y int) Field {
	fields := make([][]int64, y)
	for i := 0; i < y; i++ {
		fields[i] = make([]int64, x)
	}

	c := Field{
		Vents:      fields,
		readyToUse: true,
	}
	return c
}
func (f *Field) AddLines(lines []string, diagonal bool) {
	for _, line := range lines {
		f.addLine(line, diagonal)
	}
}

func (f *Field) addLine(line string, diagonal bool) {
	fmt.Println(line)
	x1, y1, x2, y2, dir := extractCoordinates(line)
	switch dir {
	case "diagonal":
		if diagonal {
			f.addDiagonalLine(x1, y1, x2, y2)
		}
	case "vertical":
		f.addVerticalLine(x1, y1, x2, y2)
	case "horizontal":
		f.addHorizontalLine(x1, y1, x2, y2)
	}
}

func (f *Field) addVerticalLine(x1, y1, x2, y2 int) {
	ry1, ry2 := lowerFirst(y1, y2)
	fmt.Printf("Vertical %d %d %d %d\n", x1, ry1, x2, ry2)
	for y := ry1; y <= ry2; y++ {
		fmt.Printf("y %d ry1 %d ry2 %d\n", y, ry1, ry2)
		fmt.Printf("Adding %d %d\n", x1, y)
		f.Vents[x1][y]++
	}

}

func (f *Field) addHorizontalLine(x1, y1, x2, y2 int) {
	rx1, rx2 := lowerFirst(x1, x2)
	fmt.Printf("Horizontal %d %d %d %d\n", rx1, y1, rx2, y2)
	for x := rx1; x <= rx2; x++ {

		fmt.Printf("x %d rx1 %d rx2 %d\n", x, rx1, rx2)
		fmt.Printf("Adding %d %d\n", x, y1)
		f.Vents[x][y1]++
	}
}

func (f *Field) addDiagonalLine(x1, y1, x2, y2 int) {
	return
}

func (f *Field) VentCountGreaterOne() int {
	count := 0
	for y := 0; y < len(f.Vents); y++ {
		for x := 0; x < len(f.Vents[0]); x++ {
			if f.Vents[x][y] > 1 {
				count++
			}
		}
	}
	return count
}

func lowerFirst(n1, n2 int) (int, int) {
	if n1 > n2 {
		return n2, n1
	}
	return n1, n2
}

func extractCoordinates(s string) (x1, y1, x2, y2 int, dir string) {
	coords := strings.Split(s, " -> ")
	from := strings.Split(coords[0], ",")
	to := strings.Split(coords[1], ",")

	x1, _ = strconv.Atoi(from[0])
	y1, _ = strconv.Atoi(from[1])
	x2, _ = strconv.Atoi(to[0])
	y2, _ = strconv.Atoi(to[1])
	if x1 != x2 && y1 != y2 {
		dir = "diagonal"
		return
	}
	if x1 != x2 && y1 == y2 {
		dir = "horizontal"
		return
	}

	if x1 == x2 && y1 != y2 {
		dir = "vertical"
		return
	}

	dir = "bad"
	return
}
