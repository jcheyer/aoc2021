package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/pathfinder"
	"github.com/jcheyer/aoc2021/lib"
)

type Day15 struct {
	rawInput []string
}

func (d *Day15) Load(file string) error {

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

func (d *Day15) Part1() string {
	pf, err := pathfinder.New(d.rawInput, 1)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%d", pf.Best1())

}

func (d *Day15) Part2() string {
	pf, err := pathfinder.New(d.rawInput, 5)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%d", pf.Best2())
}
