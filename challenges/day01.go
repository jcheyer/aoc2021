package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/lib"
)

type Day01 struct {
	rawInput  []string
	sonarData []int64
}

func (d *Day01) Load(file string) error {
	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	d.sonarData, err = lib.ParseLines2Int(d.rawInput)
	if err != nil {
		return err
	}
	return nil
}

func (d *Day01) Part1() string {
	var prev int64 = d.sonarData[0]
	var counter int64 = 0
	for i := 1; i < len(d.sonarData); i++ {
		if d.sonarData[i] > prev {
			counter++
		}
		prev = d.sonarData[i]
	}

	return fmt.Sprintf("%d", counter)
}

func (d *Day01) Part2() string {

	var prev int64 = d.sonarData[0] + d.sonarData[1] + d.sonarData[2]
	var counter int64 = 0
	for i := 1; i < len(d.sonarData)-2; i++ {
		window := d.sonarData[i] + d.sonarData[i+1] + d.sonarData[i+2]
		if window > prev {
			counter++
		}
		prev = window
	}

	return fmt.Sprintf("%d", counter)

}
