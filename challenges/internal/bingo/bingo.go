package bingo

import (
	"strconv"
	"strings"
)

type Card struct {
	Fields      [][]int64
	readyToLoad bool
	loaded      bool
}

func New() Card {
	size := 5

	fields := make([][]int64, size)
	for i := 0; i < size; i++ {
		fields[i] = make([]int64, size)
	}

	c := Card{
		Fields:      fields,
		readyToLoad: true,
	}
	return c
}
func (c *Card) Setup(input []string) error {
	if !c.readyToLoad {
		panic("dooo")
	}

	for i, line := range input {
		numbers := strings.Fields(line)
		for j, number := range numbers {
			res, err := strconv.ParseInt(number, 10, 64)
			if err != nil {
				return err
			}
			c.Fields[i][j] = res

		}
	}
	return nil
}

func (c *Card) MarkNumber(number int64) {
	for y := 0; y < len(c.Fields); y++ {
		for x := 0; x < len(c.Fields[0]); x++ {
			if c.Fields[x][y] == number {
				c.Fields[x][y] = convert(number)
			}
		}
	}
}

func (c *Card) HasBingo() bool {

	for y := 0; y < len(c.Fields); y++ {
		bingo := true
		for x := 0; x < len(c.Fields[y]); x++ {
			if !c.isFieldSet(x, y) {
				bingo = false
			}
		}
		if bingo {
			return true
		}
	}

	for x := 0; x < len(c.Fields[0]); x++ {
		bingo := true
		for y := 0; y < len(c.Fields); y++ {
			if !c.isFieldSet(x, y) {
				bingo = false
			}
		}
		if bingo {
			return true
		}
	}

	return false
}

func (c *Card) SumNotHit() int64 {
	var sum int64 = 0
	for y := 0; y < len(c.Fields); y++ {
		for x := 0; x < len(c.Fields[y]); x++ {
			if !c.isFieldSet(x, y) {
				sum += c.Fields[x][y]
			}
		}
	}
	return sum
}

func convert(number int64) int64 {
	if number < 0 {
		return number + 1000
	}
	return number - 1000
}

func (c *Card) isFieldSet(x, y int) bool {
	return c.Fields[x][y] < 0
}
