package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/lib"
)

type Day09 struct {
	rawInput []string
}

func (d *Day09) Load(file string) error {

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

func (d *Day09) Part1() string {

	return fmt.Sprintf("%d", 0)
}

func (d *Day09) Part2() string {

	return fmt.Sprintf("%d", 0)
}
