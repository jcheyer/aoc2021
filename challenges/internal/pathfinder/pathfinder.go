package pathfinder

import (
	"strconv"

	"github.com/RyanCarrier/dijkstra"
)

type pathFinder struct {
	d            *dijkstra.Graph
	data         [][]uint8
	naturalSizeX int
	naturalSizeY int
	mult         int
	sizeX        int
	sizeY        int
}

func New(data []string, extend int) (*pathFinder, error) {
	pf := &pathFinder{
		naturalSizeX: len(data[0]),
		naturalSizeY: len(data),
		mult:         extend,
		sizeX:        len(data[0]) * extend,
		sizeY:        len(data) * extend,
	}

	arc, err := parseLines2Int(data)
	if err != nil {
		return pf, err
	}
	pf.data = arc

	pf.d = dijkstra.NewGraph()

	// Create all vertex
	for y := 0; y < pf.sizeY; y++ {
		for x := 0; x < pf.sizeX; x++ {
			vid := pf.coord2vertex(x, y)

			pf.d.AddVertex(vid)
		}
	}

	// Add all arcs
	for y := 0; y < pf.sizeY; y++ {
		for x := 0; x < pf.sizeX; x++ {
			if y < pf.sizeY-1 {
				err := pf.d.AddArc(pf.coord2vertex(x, y), pf.coord2vertex(x, y+1), pf.val(x, y+1))
				if err != nil {
					return pf, err
				}
			}

			if y != 0 {
				err := pf.d.AddArc(pf.coord2vertex(x, y), pf.coord2vertex(x, y-1), pf.val(x, y-1))
				if err != nil {
					return pf, err
				}
			}

			if x != 0 {
				err := pf.d.AddArc(pf.coord2vertex(x, y), pf.coord2vertex(x-1, y), pf.val(x-1, y))
				if err != nil {
					return pf, err
				}
			}

			if x < pf.sizeX-1 {
				err := pf.d.AddArc(pf.coord2vertex(x, y), pf.coord2vertex(x+1, y), pf.val(x+1, y))
				if err != nil {
					return pf, err
				}
			}
		}
	}

	return pf, nil
}

func (pf *pathFinder) coord2vertex(x, y int) int {
	return pf.sizeX*(y) + x
}

func (pf *pathFinder) Best1() int {
	p, _ := pf.d.Shortest(0, pf.coord2vertex(pf.naturalSizeX-1, pf.naturalSizeY-1))
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

func (pf *pathFinder) val(x, y int) int64 {
	rx := x
	multX := 0
	if x >= pf.naturalSizeX {
		multX = x / pf.naturalSizeX
		rx = x % pf.naturalSizeX
	}

	ry := y
	multY := 0
	if y >= pf.naturalSizeY {
		multY = y / pf.naturalSizeY
		ry = y % pf.naturalSizeY
	}

	val := gen(pf.data[ry][rx], multX+multY)
	return int64(val)
}
