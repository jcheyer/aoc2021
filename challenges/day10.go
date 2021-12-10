package challenges

import (
	"fmt"
	"log"

	"github.com/jcheyer/aoc2021/lib"
)

type Day10 struct {
	rawInput []string
}

func (d *Day10) Load(file string) error {

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

func (d *Day10) Part1() string {
	score := 0
	for _, line := range d.rawInput {
		score = score + checkLine(line)
	}
	return fmt.Sprintf("%d", score)
}

func (d *Day10) Part2() string {

	return fmt.Sprintf("%d", 0)
}

func checkLine(line string) int {
	s := lib.NewStack()
	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			s.Push(r)
		case ')':
			c, err := s.Pop()
			if err != nil {
				log.Panicf("stack errorr %c in %s", r, line)
			}
			if c != '(' {
				return 3
			}
		case ']':
			c, err := s.Pop()
			if err != nil {
				log.Panicf("stack errorr %c in %s", r, line)
			}
			if c != '[' {
				return 57
			}
		case '}':
			c, err := s.Pop()
			if err != nil {
				log.Panicf("stack errorr %c in %s", r, line)
			}
			if c != '{' {
				return 1197
			}
		case '>':
			c, err := s.Pop()
			if err != nil {
				log.Panicf("stack error %c in %s", r, line)
			}
			if c != '<' {
				return 25137
			}
		default:
			log.Panicf("bad char %c in %s", r, line)
		}
	}
	return 0
}
