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
	assert.Len(t, poly.Rules, 16)
	assert.Len(t, poly.poly, 3)
	h, l := poly.Do(10)
	//h, l := poly.Do(10)
	/*poly.do()
	spew.Dump(poly.poly)
	assert.Len(t, poly.poly, 6)
	spew.Dump(poly.poly)
	assert.Equal(t, poly.poly, poly.stringToPoly("NCNBCHB"))
	poly.do()
	assert.Equal(t, poly.poly, poly.stringToPoly("NBCCNBBBCBHCB"))
	poly.do()
	assert.Equal(t, poly.poly, poly.stringToPoly("NBBBCNCCNBBNBNBBCHBHHBCHB"))
	assert.Equal(t, poly.poly, poly.stringToPoly("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB"))*/

	assert.Equal(t, h, 1749)
	assert.Equal(t, l, 161)

}
