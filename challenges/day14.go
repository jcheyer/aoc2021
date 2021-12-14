package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/polymer"
	"github.com/jcheyer/aoc2021/lib"
)

type Day14 struct {
	rawInput []string
}

func (d *Day14) Load(file string) error {

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

func (d *Day14) Part1() string {
	p := polymer.New(d.rawInput)

	h, l := p.Do(10)

	return fmt.Sprintf("%d", h-l)
}

func (d *Day14) Part2() string {
	p := polymer.New(d.rawInput)

	h, l := p.Do(40)

	return fmt.Sprintf("%d", h-l)
}
