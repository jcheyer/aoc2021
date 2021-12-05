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

	return nil
}

func (d *Day05) Part1() string {
	d.field = vents.New(1000, 1000)
	d.field.AddLines(d.rawInput, false)
	return fmt.Sprintf("%d", d.field.VentCountGreaterOne(1))
}

func (d *Day05) Part2() string {
	d.field = vents.New(1000, 1000)
	d.field.AddLines(d.rawInput, true)
	return fmt.Sprintf("%d", d.field.VentCountGreaterOne(1))
}
