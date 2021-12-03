package challenges

import (
	"github.com/jcheyer/aoc2021/lib"
)

type Day04 struct {
	rawInput []string
}

func (d *Day04) Load(file string) error {
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

func (d *Day04) Part1() string {
	return ""
}

func (d *Day04) Part2() string {

	return ""
}
