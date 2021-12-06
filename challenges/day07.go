package challenges

import (
	"github.com/jcheyer/aoc2021/lib"
)

type Day07 struct {
	rawInput []string
}

func (d *Day07) Load(file string) error {

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

func (d *Day07) Part1() string {
	return ""
}

func (d *Day07) Part2() string {
	return ""
}
