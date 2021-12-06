package fishes

import (
	"strconv"
	"strings"
)

type BetterLantern struct {
	Fishes    []uint64
	ready     bool
	lifecycle int
}

func NewBetterLantern(s string) *BetterLantern {
	sworm := BetterLantern{
		Fishes: make([]uint64, 9),
	}
	fishes := strings.Split(s, ",")
	for _, fish := range fishes {
		lifecycle, _ := strconv.Atoi(fish)
		sworm.Fishes[lifecycle]++

	}

	sworm.ready = true

	return &sworm
}

func (b *BetterLantern) Grow(lifecycle int) {
	for i := 0; i < lifecycle; i++ {
		//spew.Dump(b.Fishes)
		born := b.Fishes[0]
		b.Fishes[0] = b.Fishes[1]
		b.Fishes[1] = b.Fishes[2]
		b.Fishes[2] = b.Fishes[3]
		b.Fishes[3] = b.Fishes[4]
		b.Fishes[4] = b.Fishes[5]
		b.Fishes[5] = b.Fishes[6]
		b.Fishes[6] = b.Fishes[7]
		b.Fishes[7] = b.Fishes[8]
		b.Fishes[8] = born
		b.Fishes[6] = b.Fishes[6] + born
		//spew.Dump(b.Fishes)
		//spew.Dump("==========")
		b.lifecycle++
	}
}

func (b *BetterLantern) Sum() uint64 {
	return b.Fishes[0] + b.Fishes[1] + b.Fishes[2] + b.Fishes[3] + b.Fishes[4] + b.Fishes[5] + b.Fishes[6] + b.Fishes[7] + b.Fishes[8]
}
