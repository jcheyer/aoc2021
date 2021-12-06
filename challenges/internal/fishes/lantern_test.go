package fishes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New("3,4,3,1,2")
	assert.Equal(t, &Sworm{Fishes: []uint8{0x3, 0x4, 0x3, 0x1, 0x2}, ready: true}, s)
	s.Grow(1)
	assert.Equal(t, &Sworm{Fishes: []uint8{0x2, 0x3, 0x2, 0x0, 0x1}, ready: true}, s)
	s.Grow(1)
	assert.Equal(t, &Sworm{Fishes: []uint8{0x1, 0x2, 0x1, 0x6, 0x0, 0x8}, ready: true}, s)
	s.Grow(1)
	assert.Equal(t, &Sworm{Fishes: []uint8{0x0, 0x1, 0x0, 0x5, 0x6, 0x7, 0x8}, ready: true}, s)
	s.Grow(2)
	assert.Equal(t, &Sworm{Fishes: []uint8{0x5, 0x6, 0x5, 0x3, 0x4, 0x5, 0x6, 0x7, 0x7, 0x8}, ready: true}, s)
	s.Grow(75)
	assert.Len(t, s.Fishes, 5934)
	//s.Grow(85)
	//s.Grow(100)
	//assert.Len(t, s.Fishes, 26984457539)

}
