package diracdice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	d := NewDDice(0)
	g := NewGame(4, 8, d)
	g.Turn()
	assert.Equal(t, 10, g.s1)
	assert.Equal(t, 10, g.p1)

	g.Turn()
	assert.Equal(t, 3, g.s2)
	assert.Equal(t, 3, g.p2)

	g.Turn()
	assert.Equal(t, 14, g.s1)
	assert.Equal(t, 4, g.p1)

	g.Turn()
	assert.Equal(t, 9, g.s2)
	assert.Equal(t, 6, g.p2)
}

func TestWinner(t *testing.T) {
	d := NewDDice(0)
	g := NewGame(4, 8, d)
	for g.Winner() == 0 {
		g.Turn()
	}
	assert.Equal(t, 739785, g.Score())
}
