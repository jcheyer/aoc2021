package graph

import (
	"strings"
	"unicode"
)

type linkMap map[string][]string

type caveWorld struct {
	nodes linkMap
}

func New(input []string) *caveWorld {
	cw := caveWorld{
		nodes: make(linkMap),
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

func (cw *caveWorld) Solve1() int {
	paths := findPaths(cw.nodes, []string{"start"}, make(map[string]byte), 1)

	return len(paths)
}

func (cw *caveWorld) Solve2() int {
	paths := findPaths(cw.nodes, []string{"start"}, make(map[string]byte), 2)

	return len(paths)
}

func findPaths(links linkMap, currentPath []string, visited map[string]byte, maxSingleSmallVisits byte) [][]string {
	currentCave := currentPath[len(currentPath)-1]

	if currentCave == "end" {
		return [][]string{currentPath}
	}

	visited[currentCave]++

	paths := [][]string{}
	for _, nextCave := range links[currentCave] {
		if !alreadyMaxVisits(visited, nextCave, maxSingleSmallVisits) {
			nextPaths := findPaths(links, append(currentPath, nextCave), cloneMap(visited), maxSingleSmallVisits)
			paths = append(paths, nextPaths...)
		}
	}
	return paths
}

func isSmallCave(cave string) bool {
	return !unicode.IsUpper([]rune(cave)[0])
}

func cloneMap(m map[string]byte) map[string]byte {
	r := make(map[string]byte)
	for k, v := range m {
		r[k] = v
	}
	return r
}

func alreadyMaxVisits(visited map[string]byte, cave string, maxSingleSmallVisits byte) bool {
	if !isSmallCave(cave) || visited[cave] == 0 {
		return false
	} else if cave == "start" || cave == "end" || maxSingleSmallVisits < 2 {
		return true
	}

	for k, v := range visited {
		if isSmallCave(k) && v >= maxSingleSmallVisits {
			return true
		}
	}

	return false
}

/*func addLink(links linkMap, a string, b string) {
	links[a] = append(links[a], b)
}
*/
