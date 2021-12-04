package bingo

type Card struct {
	Fields [][]int64
	loaded bool
}

func New() Card {
	size := 5

	fields := make([][]int64, size)
	for i := 0; i < size; i++ {
		fields[i] = make([]int64, size)
	}

	c := Card{Fields: fields}
	return c
}
func (c *Card) Setup(input string) {

}

func (c *Card) MarkNumber(number int64) {
	for y := 0; y < len(c.Fields); y++ {
		for x := 0; x < len(c.Fields[0]); x++ {
			if c.Fields[x][y] == number {
				c.Fields[y][x] -= convert(number)
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

func convert(number int64) int64 {
	if number < 0 {
		return number + 1000
	}
	return number - 1000
}

func (c *Card) isFieldSet(x, y int) bool {
	return c.Fields[x][y] < 0
}
