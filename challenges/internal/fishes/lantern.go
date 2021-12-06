package fishes

import (
	"strconv"
	"strings"
)

type Sworm struct {
	Fishes []byte
	ready  bool
}

func New(s string) *Sworm {
	sworm := Sworm{
		Fishes: make([]byte, 0),
	}
	fishes := strings.Split(s, ",")
	for _, fish := range fishes {
		lifecycle, _ := strconv.Atoi(fish)
		sworm.Fishes = append(sworm.Fishes, byte(lifecycle))

	}

	sworm.ready = true

	return &sworm
}

func (sworm *Sworm) Grow(days int) {
	for i := 0; i < days; i++ {
		newFishes := make([]byte, 0)
		for i, fish := range sworm.Fishes {
			switch fish {
			case 0:
				newFishes = append(newFishes, 8)
				sworm.Fishes[i] = byte(6)
			default:
				sworm.Fishes[i]--
			}
		}
		sworm.Fishes = append(sworm.Fishes, newFishes...)
	}
}
