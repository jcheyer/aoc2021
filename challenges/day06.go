package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/fishes"
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
	sworm := fishes.NewBetterLantern(d.rawInput[0])
	sworm.Grow(80)
	return fmt.Sprintf("%d", sworm.Sum())
}

func (d *Day06) Part2() string {
	sworm := fishes.NewBetterLantern(d.rawInput[0])
	sworm.Grow(256)
	return fmt.Sprintf("%d", sworm.Sum())
}
