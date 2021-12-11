package dumbooctopus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []string{
	"5483143223",
	"2745854711",
	"5264556173",
	"6141336146",
	"6357385478",
	"4167524645",
	"2176841721",
	"6882881134",
	"4846848554",
	"5283751526",
}

func TestOcto1(t *testing.T) {
	o := New(testData)
	count := o.Step()
	assert.Equal(t, 0, count)
	count += o.Step()
	assert.Equal(t, 35, count)
	count += o.Step()
	count += o.Step()
	count += o.Step()
	count += o.Step()
	count += o.Step()
	count += o.Step()
	count += o.Step()
	count += o.Step()
	assert.Equal(t, 204, count)
}

func TestOcto2(t *testing.T) {
	o := New(testData)
	count := 0
	for i := 0; i < 100; i++ {
		count += o.Step()
	}
	assert.Equal(t, 1656, count)
}
