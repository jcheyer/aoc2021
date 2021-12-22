package challenges

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jcheyer/aoc2021/lib"
)

type Day22 struct {
	rawInput []string
}

func (d *Day22) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	return nil
}

type coord struct {
	x int
	y int
	z int
}

func (d *Day22) Part1() string {
	var re = regexp.MustCompile(`(?m)^(on|off).x=(-{0,1}\d+)..(-{0,1}\d+),y=(-{0,1}\d+)..(-{0,1}\d+),z=(-{0,1}\d+)..(-{0,1}\d+)`)
	reactor := make(map[coord]bool)
	for _, l := range d.rawInput {
		match := re.FindStringSubmatch(l)
		onoff := match[1]
		x1, _ := strconv.Atoi(match[2])
		x2, _ := strconv.Atoi(match[3])
		y1, _ := strconv.Atoi(match[4])
		y2, _ := strconv.Atoi(match[5])
		z1, _ := strconv.Atoi(match[6])
		z2, _ := strconv.Atoi(match[7])
		if !between(x1, -50, 50) || !between(x2, -50, 50) ||
			!between(y1, -50, 50) || !between(y2, -50, 50) ||
			!between(z1, -50, 50) || !between(z2, -50, 50) {
			continue
		}
		fmt.Printf("%s\n", l)
		fmt.Printf("%s %d..%d %d..%d %d..%d\n", onoff, x1, x2, y1, y2, z1, z2)
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				for z := z1; z <= z2; z++ {
					reactor[coord{x, y, z}] = onoff == "on"
				}
			}
		}
	}
	c := 0
	for _, b := range reactor {
		if b {
			c++
		}
	}
	return fmt.Sprintf("%d", c)
}

func (d *Day22) Part2() string {
	return "No way....."
}

func between(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}
