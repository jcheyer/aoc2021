package challenges

import (
	"github.com/jcheyer/aoc2021/lib"
)

type Day19 struct {
	rawInput []string
}

func (d *Day19) Load(file string) error {

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

func (d *Day19) Part1() string {
	return "0"
}

func (d *Day19) Part2() string {
	return "0"
}
