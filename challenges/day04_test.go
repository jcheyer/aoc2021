package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay04(t *testing.T) {
	d4 := Day04{}
	err := d4.Load("testfiles/d04demo.txt")
	assert.NoError(t, err)

	assert.Len(t, d4.numbers, 27)
}
