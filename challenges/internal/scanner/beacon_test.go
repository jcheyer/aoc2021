package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBeacon(t *testing.T) {
	b, err := NewBeacon("404,-588,-901")
	assert.NoError(t, err)
	c := coord{
		x: 404, y: -588, z: -901}
	expb := beacon{coord: c}
	assert.Equal(t, expb, b)
}
