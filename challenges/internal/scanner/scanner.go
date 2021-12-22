package scanner

import (
	"github.com/jcheyer/aoc2021/lib"
)

type coord struct {
	x, y, z int
}

type scanner struct {
	ident   string
	root    bool
	beacons []beacon
}

func New(s []string) *[]scanner {
	chunks := lib.Chunks(s)
	scanners := make([]scanner, 0)

	i := 1
	for i < len(chunks) {
		sc := newScanner(chunks[i])
		scanners = append(scanners, sc)
	}

	scanners[0].root = true
	return &scanners
}

func newScanner(s []string) scanner {
	sc := scanner{
		ident:   s[0],
		beacons: make([]beacon, 0),
	}
	i := 1
	for i < len(s) {
		b, err := NewBeacon(s[i])
		if err != nil {
			panic(err)
		}
		sc.beacons = append(sc.beacons, b)
	}
	return sc
}
