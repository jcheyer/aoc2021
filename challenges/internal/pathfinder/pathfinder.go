package pathfinder

import (
	"strconv"

	"github.com/RyanCarrier/dijkstra"
)

type pathFinder struct {
	d     *dijkstra.Graph
	data  [][]uint8
	mult  int
	sizeX int
	sizeY int
	planX int
	planY int
}

func New(data []string, extend int) (*pathFinder, error) {

	pf := &pathFinder{
		mult:  extend,
		sizeX: len(data[0]) * extend,
		sizeY: len(data) * extend,
		planX: len(data[0]),
		planY: len(data),
	}

	arc, err := parseLines2Int(data)
	if err != nil {
		return pf, err
	}
	pf.data = arc

	pf.d = dijkstra.NewGraph()

	for y := 0; y < pf.sizeY; y++ {
		for x := 0; x < pf.sizeX; x++ {
			vid := pf.coord2vertex(x, y)

			pf.d.AddVertex(vid)
		}
	}

	for y := 0; y < pf.sizeY; y++ {
		for x := 0; x < pf.sizeX; x++ {
			if y < pf.sizeY-1 {
				err := pf.d.AddArc(pf.coord2vertex(x, y), pf.coord2vertex(x, y+1), int64(val(arc, x, y+1)))
				if err != nil {
					return pf, err
				}
			}

			if y != 0 {
				err := pf.d.AddArc(pf.coord2vertex(x, y), pf.coord2vertex(x, y-1), int64(val(arc, x, y-1)))
				if err != nil {
					return pf, err
				}
			}

			if x != 0 {
				err := pf.d.AddArc(pf.coord2vertex(x, y), pf.coord2vertex(x-1, y), int64(val(arc, x-1, y)))
				if err != nil {
					return pf, err
				}
			}

			if x < pf.sizeX-1 {
				err := pf.d.AddArc(pf.coord2vertex(x, y), pf.coord2vertex(x+1, y), int64(val(arc, x+1, y)))
				if err != nil {
					return pf, err
				}
			}
		}
	}

	return pf, nil
}

func (pf *pathFinder) coord2vertex(x, y int) int {
	//return fmt.Sprintf("%d-%d", x, y)

	return pf.sizeX*(y) + x
}

func (pf *pathFinder) Best1() int {
	p, _ := pf.d.Shortest(0, pf.coord2vertex(pf.planX-1, pf.planY-1))
	return int(p.Distance)
}

func (pf *pathFinder) Best2() int {
	p, _ := pf.d.Shortest(0, pf.coord2vertex(pf.sizeX-1, pf.sizeY-1))
	return int(p.Distance)
}

func parseLines2Int(lines []string) ([][]uint8, error) {
	width := len(lines[0])
	field := make([][]uint8, 0)

	for _, line := range lines {
		row := make([]uint8, width)
		for i, r := range line {
			x, err := strconv.Atoi(string(r))
			if err != nil {
				return field, err
			}

			row[i] = uint8(x)
			if err != nil {
				return field, err
			}
		}

		field = append(field, row)
	}

	return field, nil
}

func gen(i uint8, count int) uint8 {
	res := i + uint8(count)
	if res > 9 {
		res = res - 9
	}
	return res
}

func val(arc [][]uint8, x, y int) uint8 {
	rx := x
	multX := 0
	if x >= len(arc[0]) {
		multX = x / len(arc[0])
		rx = x % len(arc[0])
	}
	ry := y
	multY := 0
	if y >= len(arc) {
		multY = y / len(arc)
		ry = y % len(arc)
	}
	//fmt.Printf("x: %d y: %d rx:%d ry: %d multX: %d multY: %d\n", x, y, rx, ry, multX, multY)
	val := gen(arc[ry][rx], multX+multY)
	return val
}
