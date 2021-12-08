package challenges

import (
	"fmt"
	"strings"

	"github.com/jcheyer/aoc2021/challenges/internal/sevensegments"
	"github.com/jcheyer/aoc2021/lib"
)

type Day08 struct {
	rawInput []string
}

func (d *Day08) Load(file string) error {

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

func (d *Day08) Part1() string {
	var count uint64
	for _, line := range d.rawInput {
		parts := strings.Split(line, "|")
		outputs := strings.Fields(parts[1])

		for _, output := range outputs {
			l := len(output)
			switch l {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	return fmt.Sprintf("%d", count)
}

func (d *Day08) Part2() string {
	result := 0
	for _, line := range d.rawInput {
		s := sevensegments.New(line)
		result = result + s.Solve()

	}

	return fmt.Sprintf("%d", result)

}
