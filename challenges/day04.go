package challenges

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jcheyer/aoc2021/challenges/internal/bingo"
	"github.com/jcheyer/aoc2021/lib"
)

type Day04 struct {
	rawInput []string
	numbers  []int64
	cards    []bingo.Card
	cardsWon []bool
	lastDraw int
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

	lineCount := len(d.rawInput)
	cardsCount := (lineCount - 1) / 5

	for i := 0; i < cardsCount; i++ {
		a := 1 + i*5
		b := 6 + i*5
		data := d.rawInput[a:b]

		c := bingo.New()
		err := c.Setup(data)
		if err != nil {
			return err
		}
		d.cards = append(d.cards, c)
	}

	d.cardsWon = make([]bool, len(d.cards))

	return nil
}

func (d *Day04) Part1() string {
	for iDraw, draw := range d.numbers {
		d.lastDraw = iDraw
		for iCard, card := range d.cards {
			card.MarkNumber(draw)
			if card.HasBingo() {
				d.cardsWon[iCard] = true
				sum := card.SumNotHit()

				return fmt.Sprintf("%d", sum*draw)

			}
		}
	}
	panic("Part1 not working")
}

func (d *Day04) Part2() string {

	for iDraw := d.lastDraw; iDraw < len(d.numbers); iDraw++ {
		draw := d.numbers[iDraw]
		for i, card := range d.cards {
			card.MarkNumber(draw)
			if card.HasBingo() {
				if !d.cardsWon[i] {
					d.cardsWon[i] = true
					nrOfBoardsWon := 0
					for _, hasWon := range d.cardsWon {
						if hasWon {
							nrOfBoardsWon++
						}
					}

					if nrOfBoardsWon == len(d.cards) {
						sum := card.SumNotHit()
						return fmt.Sprintf("%d", sum*draw)

					}
				}
			}
		}
	}
	return "Part2 not working"
}
