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
	//fmt.Println(line)
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
	for y := ry1; y <= ry2; y++ {
		f.Vents[x1][y]++
	}

}

func (f *Field) addHorizontalLine(x1, y1, x2, y2 int) {
	rx1, rx2 := lowerFirst(x1, x2)
	for x := rx1; x <= rx2; x++ {
		f.Vents[x][y1]++
	}
}

func (f *Field) addDiagonalLine(x1, y1, x2, y2 int) {
	count := x1 - x2
	if count < 0 {
		count = count * -1
	}

	xInc := -1
	if x1 < x2 {
		xInc = 1
	}

	yInc := -1
	if y1 < y2 {
		yInc = 1
	}

	for i := 0; i <= count; i++ {
		f.Vents[x1+i*xInc][y1+i*yInc]++
	}
}

func (f *Field) VentCountGreaterOne(limit int64) int {
	count := 0
	for y := 0; y < len(f.Vents); y++ {
		for x := 0; x < len(f.Vents[0]); x++ {
			if f.Vents[x][y] > limit {
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

func (f *Field) String() string {
	var sb strings.Builder
	sb.WriteString("\n")
	for y := 0; y < len(f.Vents); y++ {
		for x := 0; x < len(f.Vents[0]); x++ {
			if f.Vents[x][y] > 0 {
				sb.WriteString(fmt.Sprintf("%d ", f.Vents[x][y]))
			} else {
				sb.WriteString(". ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
