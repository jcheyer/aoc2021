package challenges

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD7(t *testing.T) {

	d7 := Day07{}
	err := d7.Load("testfiles/d07demo.txt")
	assert.NoError(t, err)
	result := d7.Part1()
	assert.Equal(t, "2", result)
}
