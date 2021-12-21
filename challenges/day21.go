package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/diracdice"
	"github.com/jcheyer/aoc2021/lib"
)

type Day21 struct {
	rawInput []string
}

func (d *Day21) Load(file string) error {

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

func (d *Day21) Part1() string {
	//Player 1 starting position: 8
	//Player 2 starting position: 6
	dice := diracdice.NewDDice(0)
	g := diracdice.NewGame(8, 6, dice)
	for g.Winner() == 0 {
		g.Turn()
	}
	return fmt.Sprintf("%d", g.Score())

}

func (d *Day21) Part2() string {
	return "0"
}
