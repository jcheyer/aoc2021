package folding

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jcheyer/aoc2021/lib"
)

type coord struct {
	x, y int
}

type paper struct {
	dots  map[coord]bool
	Folds []coord
}

func New(data []string) *paper {
	p := paper{
		dots:  make(map[coord]bool, 0),
		Folds: make([]coord, 0),
	}

	regCoord := regexp.MustCompile(`^(\d+),(\d+)$`)
	regFoldX := regexp.MustCompile(`^fold along x=(\d+)$`)
	regFoldY := regexp.MustCompile(`^fold along y=(\d+)$`)

	for _, line := range data {
		if match := regCoord.FindStringSubmatch(line); match != nil {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])

			//p.dots = append(p.dots, coord{x: x, y: y})
			p.dots[coord{x: x, y: y}] = true
		}

		if match := regFoldX.FindStringSubmatch(line); match != nil {
			x, _ := strconv.Atoi(match[1])

			p.Folds = append(p.Folds, coord{x: x, y: 0})

		}

		if match := regFoldY.FindStringSubmatch(line); match != nil {
			y, _ := strconv.Atoi(match[1])

			p.Folds = append(p.Folds, coord{x: 0, y: y})
		}
	}

	return &p
}

func (p *paper) CountDots() int {
	return len(p.dots)
}

func (p *paper) Fold(fold coord) {
	if fold.x == 0 {
		p.foldY(fold.y)
	} else {
		p.foldX(fold.x)
	}
}

func (p *paper) foldX(x int) {
	doubeX := 2 * x
	for k, _ := range p.dots {
		if k.x > x {
			newX := doubeX - k.x
			p.dots[coord{x: newX, y: k.y}] = true
			delete(p.dots, k)
		}
	}

}

func (p *paper) foldY(y int) {

	doubeY := 2 * y
	for k, _ := range p.dots {
		if k.y > y {
			newY := doubeY - k.y
			p.dots[coord{x: k.x, y: newY}] = true
			delete(p.dots, k)
		}
	}

}

func (p *paper) String() string {
	var sizeX, sizeY int

	for k, _ := range p.dots {
		sizeX = lib.HighInt(sizeX, k.x)
		sizeY = lib.HighInt(sizeY, k.y)
	}

	var sb strings.Builder
	for y := 0; y <= sizeY; y++ {
		for x := 0; x <= sizeX; x++ {
			switch p.dots[coord{x: x, y: y}] {
			case true:
				sb.WriteString("#")
			case false:
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString("\n")
	return sb.String()
}
