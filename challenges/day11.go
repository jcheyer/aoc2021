package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/dumbooctopus"
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
	o := dumbooctopus.New(d.rawInput)
	count := 0
	for i := 0; i < 100; i++ {
		count += o.Step()
	}
	return fmt.Sprintf("%d", count)
}

func (d *Day11) Part2() string {
	o := dumbooctopus.New(d.rawInput)
	count := 0
	for o.Sum() != 0 {
		o.Step()
		count++
	}
	return fmt.Sprintf("%d", count)
}
