package challenges

import (
	"github.com/jcheyer/aoc2021/lib"
)

type Day05 struct {
	rawInput []string
}

func (d *Day05) Load(file string) error {

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

func (d *Day05) Part1() string {
	panic("Part1 not working")
}

func (d *Day05) Part2() string {
	return "Part2 not working"
}
