package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/vents"
	"github.com/jcheyer/aoc2021/lib"
)

type Day05 struct {
	rawInput []string
	field    vents.Field
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

	d.field = vents.New(1000, 1000)

	return nil
}

func (d *Day05) Part1() string {
	d.field.AddLines(d.rawInput, false)
	return fmt.Sprintf("%d", d.field.VentCountGreaterOne())
}

func (d *Day05) Part2() string {
	return "Part2 not working"
}
