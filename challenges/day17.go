package challenges

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jcheyer/aoc2021/lib"
)

type Probe struct {
	xVel       int
	yVel       int
	xPos       int
	yPos       int
	highest    int
	hit        bool
	overTarget bool
}

type Target struct {
	xStart int
	xEnd   int
	yStart int
	yEnd   int
}
type Day17 struct {
	rawInput []string
	target   Target
}

func (d *Day17) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLines(file)
	if err != nil {
		return err
	}

	s := strings.Replace(d.rawInput[0], ",", "", 1)
	v := strings.Fields(s)
	xs := strings.Split(v[2][2:], "..")
	ys := strings.Split(v[3][2:], "..")

	xStart, err := strconv.Atoi(xs[0])
	if err != nil {
		return err
	}
	xEnd, err := strconv.Atoi(xs[1])
	if err != nil {
		return err
	}
	yStart, err := strconv.Atoi(ys[0])
	if err != nil {
		return err
	}
	yEnd, err := strconv.Atoi(ys[1])
	if err != nil {
		return err
	}

	d.target = Target{
		xStart: xStart,
		xEnd:   xEnd,
		yStart: yStart,
		yEnd:   yEnd,
	}

	return nil
}

func (d *Day17) Part1() string {
	r, _ := d.bruteForce()
	return fmt.Sprintf("%d", r)
}

func (d *Day17) Part2() string {
	_, r := d.bruteForce()
	return fmt.Sprintf("%d", r)
}

func (p *Probe) fly() {
	p.xPos += p.xVel
	p.yPos += p.yVel
	if p.xVel > 0 {
		p.xVel--
	} else if p.xVel < 0 {
		p.xVel++
	}
	p.yVel--
	if p.yPos > p.highest {
		p.highest = p.yPos
	}

}

func (p *Probe) check(t Target) {
	// HIT
	if p.xPos >= t.xStart && p.xPos <= t.xEnd {
		if p.yPos >= t.yStart && p.yPos <= t.yEnd {
			p.hit = true
		}
	}

	// MISS
	if p.xPos > t.xEnd || p.yPos < t.yStart {
		p.overTarget = true
	}
}

func (d *Day17) shoot(x, y int) (bool, int) {
	p := Probe{
		xVel: x,
		yVel: y,
	}

	for !p.hit {
		p.fly()
		p.check(d.target)
		if p.overTarget {
			break
		}
	}
	return p.hit, p.highest
}

func (d *Day17) bruteForce() (int, int) {
	highest := 0
	hits := 0
	for x := 1; x <= d.target.xEnd; x++ {
		for y := d.target.yStart; y < 15*(d.target.yEnd-d.target.yStart); y++ {

			if hit, high := d.shoot(x, y); hit {
				if high > highest {
					highest = high
				}
				hits++
			}
		}
	}
	return highest, hits
}
