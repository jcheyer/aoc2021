package challenges

import (
	"github.com/jcheyer/aoc2021/lib"
)

type Day06 struct {
	rawInput []string
}

func (d *Day06) Load(file string) error {

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

func (d *Day06) Part1() string {
	return ""
}

func (d *Day06) Part2() string {
	return ""
}
