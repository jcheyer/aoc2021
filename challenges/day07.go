package challenges

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jcheyer/aoc2021/lib"
)

type Day07 struct {
	rawInput       []string
	pos            []uint64
	maxPos, minPos uint64
}

func (d *Day07) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	v := strings.Split(d.rawInput[0], ",")
	d.pos = make([]uint64, len(v))
	for i, pos := range v {
		d.pos[i], err = strconv.ParseUint(pos, 10, 64)
		if err != nil {
			return err
		}
	}

	for _, pos := range d.pos {
		if pos < d.minPos {
			d.minPos = pos
		}
		if pos > d.maxPos {
			d.maxPos = pos
		}
	}

	return nil
}

func (d *Day07) Part1() string {
	var bestFuel uint64 = math.MaxUint64
	var bestFuelPos uint64
	var fueldiff uint64

	for pos := d.minPos; pos <= d.maxPos; pos++ {
		fueldiff = 0
		for _, checkPos := range d.pos {
			fueldiff += lib.AbsDiffUInt64(checkPos, pos)
		}
		if fueldiff < bestFuel {
			bestFuel = fueldiff
			bestFuelPos = pos
		}
	}
	return fmt.Sprintf("%d", bestFuelPos)
}

func (d *Day07) Part2() string {
	return ""
}
