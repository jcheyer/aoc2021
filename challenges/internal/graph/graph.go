package graph

import (
	"strings"
)

type caveWorld struct {
	nodes map[string][]string
	paths [][]string
}

func New(input []string) *caveWorld {
	cw := caveWorld{
		nodes: make(map[string][]string),
		paths: make([][]string, 0),
	}
	for _, l := range input {
		row := strings.Split(l, "-")

		from := row[0]
		to := row[1]
		if _, ok := cw.nodes[from]; !ok {
			cw.nodes[from] = []string{}
		}
		if _, ok := cw.nodes[to]; !ok {
			cw.nodes[to] = []string{}
		}
		cw.nodes[from] = append(cw.nodes[from], to)
		cw.nodes[to] = append(cw.nodes[to], from)
	}

	return &cw
}

func (cw *caveWorld) CountPossiblePaths() int {
	return 0
}
