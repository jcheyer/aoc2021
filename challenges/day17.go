package challenges

import (
	"github.com/jcheyer/aoc2021/lib"
)

type Day17 struct {
	rawInput []string
}

func (d *Day17) Load(file string) error {

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

func (d *Day17) Part1() string {
	return "0"
}

func (d *Day17) Part2() string {
	return "0"
}
