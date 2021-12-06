package fishes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBetter(t *testing.T) {
	s := NewBetterLantern("3,4,3,1,2")
	s.Grow(1)
	assert.Equal(t, uint64(5), s.Sum())
	s.Grow(1)
	assert.Equal(t, uint64(6), s.Sum())
	s.Grow(1)
	assert.Equal(t, uint64(7), s.Sum())
	s.Grow(1)
	assert.Equal(t, uint64(9), s.Sum())
	s.Grow(1)
	assert.Equal(t, uint64(10), s.Sum())
	s.Grow(1)
	assert.Equal(t, uint64(10), s.Sum())
	s.Grow(1)
	assert.Equal(t, uint64(10), s.Sum())
	s.Grow(1)
	assert.Equal(t, uint64(10), s.Sum())
	s.Grow(1)
	assert.Equal(t, uint64(11), s.Sum())
	s.Grow(1)
	assert.Equal(t, uint64(12), s.Sum())
	s.Grow(1)
	assert.Equal(t, uint64(15), s.Sum())
	assert.Equal(t, 11, s.lifecycle)
	s.Grow(1)
	assert.Equal(t, uint64(17), s.Sum())
	assert.Equal(t, 12, s.lifecycle)
	s.Grow(68)
	assert.Equal(t, uint64(5934), s.Sum())
	assert.Equal(t, 80, s.lifecycle)

}
