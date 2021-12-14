package polymer

import (
	"testing"

	"github.com/jcheyer/aoc2021/lib"
	"github.com/stretchr/testify/assert"
)

func TestPolymerizer(t *testing.T) {
	data, err := lib.ReadInputLines("./../../testfiles/d14demo.txt")
	assert.NoError(t, err)
	poly := New(data)
	assert.Equal(t, "NNCB", poly.Base)
	assert.Len(t, poly.rules, 16)
	assert.Equal(t, "NCNBCHB", poly.Do("NNCB"))
	assert.Equal(t, "NBCCNBBBCBHCB", poly.Do("NCNBCHB"))
	s := "NNCB"
	for i := 0; i < 10; i++ {
		s = poly.Do(s)
	}
	assert.Len(t, s, 3073)
	h, l := HighLow(ComponentCount(s))
	assert.Equal(t, h, 1749)
	assert.Equal(t, l, 161)

	for i := 0; i < 30; i++ {
		s = poly.Do(s)
	}
	h, l = HighLow(ComponentCount(s))
	assert.Equal(t, h, 2192039569602)
	assert.Equal(t, l, 3849876073)

}
