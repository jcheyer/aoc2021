package sevensegments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")
	assert.Len(t, s.digits, 10)
	assert.Len(t, s.value, 4)
	assert.Equal(t, "ab", s.digits[9])
	assert.Equal(t, 5353, s.Solve())
}

func TestSorted(t *testing.T) {
	assert.Equal(t, "abc", sort("abc"))
	assert.Equal(t, "abc", sort("cab"))
	assert.Equal(t, "abc", sort("bac"))
}

func TestStrDiff(t *testing.T) {
	assert.Equal(t, "a", strDiff("abc", "bc"))
	assert.Equal(t, "a", strDiff("bc", "abc"))
	assert.Equal(t, "ad", strDiff("bc", "abcd"))
}
