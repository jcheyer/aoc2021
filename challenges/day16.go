package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/packetdecoder"
	"github.com/jcheyer/aoc2021/lib"
)

type Day16 struct {
	rawInput []string
}

func (d *Day16) Load(file string) error {

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

func (d *Day16) Part1() string {
	decoder := packetdecoder.New(d.rawInput[0])

	packet, _ := decoder.Parse(0)
	return fmt.Sprintf("%d", packet.Versions())
}

func (d *Day16) Part2() string {
	decoder := packetdecoder.New(d.rawInput[0])

	packet, _ := decoder.Parse(0)
	return fmt.Sprintf("%+v", packet.Eval())
}
