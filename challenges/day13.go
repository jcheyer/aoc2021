package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/folding"
	"github.com/jcheyer/aoc2021/lib"
)

type Day13 struct {
	rawInput []string
}

func (d *Day13) Load(file string) error {

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

func (d *Day13) Part1() string {
	paper := folding.New(d.rawInput)
	paper.Fold(paper.Folds[0])
	return fmt.Sprintf("%d", paper.CountDots())

}

func (d *Day13) Part2() string {
	paper := folding.New(d.rawInput)
	for _, fold := range paper.Folds {
		paper.Fold(fold)
	}
	return "\n" + paper.String()
}
