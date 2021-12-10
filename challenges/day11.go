package challenges

import (
	"github.com/jcheyer/aoc2021/lib"
)

type Day11 struct {
	rawInput []string
}

func (d *Day11) Load(file string) error {

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

func (d *Day11) Part1() string {
	return "0"
}

func (d *Day11) Part2() string {
	return "0"
}
