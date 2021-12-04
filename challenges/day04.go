package challenges

import (
	"strconv"
	"strings"

	"github.com/jcheyer/aoc2021/challenges/internal/bingo"
	"github.com/jcheyer/aoc2021/lib"
)

type Day04 struct {
	rawInput []string
	numbers  []int64
	cards    []bingo.Card
}

func (d *Day04) Load(file string) error {

	d.cards = make([]bingo.Card, 0)

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	numbers := strings.Split(d.rawInput[0], ",")

	d.numbers = make([]int64, len(numbers))
	for i, n := range numbers {
		res, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			panic(err)
		}
		d.numbers[i] = res
	}

	return nil
}

func (d *Day04) Part1() string {
	return ""
}

func (d *Day04) Part2() string {

	return ""
}
