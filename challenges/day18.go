package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/snailfishmath"
	"github.com/jcheyer/aoc2021/lib"
)

type Day18 struct {
	rawInput []string
}

func (d *Day18) Load(file string) error {

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

func (d *Day18) Part1() string {
	n := snailfishmath.New(d.rawInput[0])
	for i := 1; i < len(d.rawInput); i++ {
		add := snailfishmath.New(d.rawInput[i])
		n = n.Add(add)
	}
	return fmt.Sprintf("%d", n.Magnitude())
}

func (d *Day18) Part2() string {
	return "0"
}
