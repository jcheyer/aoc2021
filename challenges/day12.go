package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/graph"
	"github.com/jcheyer/aoc2021/lib"
)

type Day12 struct {
	rawInput []string
}

func (d *Day12) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	return nil
}

func (d *Day12) Part1() string {
	cw := graph.New(d.rawInput)

	return fmt.Sprintf("%d", cw.Solve1())
}

func (d *Day12) Part2() string {
	cw := graph.New(d.rawInput)

	return fmt.Sprintf("%d", cw.Solve2())
}
