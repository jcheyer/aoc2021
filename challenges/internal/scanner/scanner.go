package scanner

import "fmt"

type coord struct {
	x, y, z int
}

type scanner struct {
	ident   string
	root    bool
	absPos  coord
	beacons []beacon
}

func New(s []string) []scanner {
	scanners := make([]scanner, 0)

	i := 0
	for i < len(s) {

		sc := scanner{
			beacons: make([]beacon, 0),
			ident:   s[i],
		}
		fmt.Printf("i: %d New Scanner  %+v\n", i, sc)
		i++ // scannerID
		for ; i < len(s); i++ {
			if s[i] == "" {
				fmt.Printf("Scanner %s done\n", sc.ident)
				i++
				scanners = append(scanners, sc)

				break

			}
			b, err := NewBeacon(s[i])
			if err != nil {
				panic(err)
			}
			sc.beacons = append(sc.beacons, b)
			fmt.Printf("i: %d beacon added %+v\n", i, sc)
		}

	}

	return scanners
}
