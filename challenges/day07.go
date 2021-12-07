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
	var fueldiff uint64

	for pos := d.minPos; pos <= d.maxPos; pos++ {

		fueldiff = d.calcFuelP1(pos)
		if fueldiff < bestFuel {
			bestFuel = fueldiff
		}
		//fmt.Printf("Pos: %d Fuel: %d\n", pos, fueldiff)
	}
	return fmt.Sprintf("%d", bestFuel)
}

func (d *Day07) calcFuelP1(x uint64) uint64 {
	var fuel uint64
	for _, checkPos := range d.pos {
		fuel += lib.AbsDiffUInt64(checkPos, x)
	}
	return fuel
}

func (d *Day07) calcFuelP2(x uint64) uint64 {
	var fuel uint64
	for _, checkPos := range d.pos {
		dist := lib.AbsDiffUInt64(checkPos, x)
		fuel += (dist*dist + dist) / 2
	}
	return fuel
}

func (d *Day07) Part2() string {
	var bestFuel uint64 = math.MaxUint64
	var fueldiff uint64

	for pos := d.minPos; pos <= d.maxPos; pos++ {

		fueldiff = d.calcFuelP2(pos)
		if fueldiff < bestFuel {
			bestFuel = fueldiff
		}
		//fmt.Printf("Pos: %d Fuel: %d\n", pos, fueldiff)
	}
	return fmt.Sprintf("%d", bestFuel)
}
