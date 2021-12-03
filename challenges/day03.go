package challenges

import (
	"fmt"
	"strconv"

	"github.com/jcheyer/aoc2021/lib"
)

type Day03 struct {
	rawInput    []string
	reportLen   int
	on          []int
	reportCount int
	gamma       string
	epsilon     string
}

func (d *Day03) Load(file string) error {
	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	d.reportLen = len(d.rawInput[0])
	d.reportCount = len(d.rawInput)

	d.on = d.countOn(d.rawInput)

	return nil
}

func (d *Day03) Part1() string {
	if len(d.rawInput) == 0 {
		return ""
	}
	d.gamma = ""
	d.epsilon = ""
	for i := 0; i < len(d.on); i++ {
		if d.reportCount-d.on[i] < d.on[i] {
			d.gamma += "1"
			d.epsilon += "0"
		} else {
			d.gamma += "0"
			d.epsilon += "1"
		}
	}

	gammaInt, _ := strconv.ParseInt(d.gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(d.epsilon, 2, 64)

	return fmt.Sprintf("%d", gammaInt*epsilonInt)
}

// externaly "inspired" (https://github.com/Alex-Wauters/advent-go-2021/blob/master/day03/binary.go)
func (d *Day03) Part2() string {
	if len(d.rawInput) == 0 {
		return ""
	}

	oxygen := d.filter(d.rawInput, 0, true)
	scrubber := d.filter(d.rawInput, 0, false)
	oxygenInt, _ := strconv.ParseInt(oxygen, 2, 64)
	scrubberInt, _ := strconv.ParseInt(scrubber, 2, 64)
	return fmt.Sprintf("%d", oxygenInt*scrubberInt)
}

func (d *Day03) countOn(input []string) []int {
	on := make([]int, len(input[0]))
	for _, line := range input {
		for i := 0; i < len(line); i++ {
			if line[i] == '1' {
				on[i]++
			}
		}

	}
	return on
}

func (d *Day03) mostAndLeastCommonRuneAtPos(lines []string, i int) (uint8, uint8) {
	count0, count1 := 0, 0
	for _, line := range lines {
		if line[i] == '0' {
			count0++
		} else {
			count1++
		}
	}
	if count0 > count1 {
		return '0', '1'
	}
	return '1', '0'
}

func (d *Day03) filter(lines []string, i int, mostCommon bool) string {
	if len(lines) == 1 {
		return lines[0]
	}
	most, least := d.mostAndLeastCommonRuneAtPos(lines, i)
	comparator := least
	if mostCommon {
		comparator = most
	}
	filtered := make([]string, 0)
	for _, l := range lines {
		if l[i] == comparator {
			filtered = append(filtered, l)
		}
	}
	return d.filter(filtered, i+1, mostCommon)
}
