package diracdice

import "github.com/jcheyer/aoc2021/lib"

type game struct {
	p1, p2   int
	s1, s2   int
	turns    int
	lastRoll int
	dice     dice
}

func NewGame(p1, p2 int, d dice) *game {
	g := &game{
		p1:   p1,
		p2:   p2,
		dice: d,
	}
	return g
}

func (g *game) Turn() {
	g.turns++
	roll := g.dice.Roll() + g.dice.Roll() + g.dice.Roll()
	g.lastRoll = roll
	switch g.turns % 2 {
	case 1: // Player 1
		g.p1 = g.p1 + roll
		g.p1 = g.p1 % 10
		if g.p1 == 0 {
			g.p1 = 10
		}
		g.s1 = g.s1 + g.p1
	case 0:
		// Player 2
		g.p2 = g.p2 + roll
		g.p2 = g.p2 % 10
		if g.p2 == 0 {
			g.p2 = 10
		}
		g.s2 = g.s2 + g.p2
	}
}

func (g *game) Winner() int {
	if g.s1 >= 1000 {
		return 1
	}
	if g.s2 > 1000 {
		return 2
	}
	return 0
}

func (g *game) Score() int {
	low := lib.LowInt(g.s1, g.s2)
	high := lib.HighInt(g.s1, g.s2)
	if high < 1000 {
		return 0
	}
	return low * g.dice.UseCount()
}
