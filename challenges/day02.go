package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/lib"
)

type Day02 struct {
	rawInput []string
}

func (d *Day02) Load(file string) error {
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

func (d *Day02) Part1() string {
	hor, depth := lib.ParseLines2Nav1(d.rawInput)

	return fmt.Sprintf("%d", hor*depth)
}

func (d *Day02) Part2() string {
	hor, depth := lib.ParseLines2Nav2(d.rawInput)

	return fmt.Sprintf("%d", hor*depth)

}
