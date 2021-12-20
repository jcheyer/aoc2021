package scanner

import (
	"strconv"
	"strings"
)

type beacon struct {
	coord
}

func NewBeacon(s string) (beacon, error) {
	b := beacon{}
	parts := strings.Split(s, ",")
	x, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return b, err
	}
	y, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return b, err
	}
	z, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return b, err
	}

	coord := coord{x: int(x), y: int(y), z: int(z)}
	b.coord = coord

	return b, nil
}
