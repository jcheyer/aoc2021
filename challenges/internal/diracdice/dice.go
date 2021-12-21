package diracdice

type dice interface {
	Roll() int
	UseCount() int
}

type ddice struct {
	lastDraw int
	count    int
}

func NewDDice(i int) *ddice {
	dd := &ddice{
		lastDraw: i,
	}
	return dd
}

func (dd *ddice) Roll() int {
	dd.lastDraw++
	if dd.lastDraw == 101 {
		dd.lastDraw = 1
	}
	dd.count++
	return dd.lastDraw
}

func (dd *ddice) UseCount() int {
	return dd.count
}
