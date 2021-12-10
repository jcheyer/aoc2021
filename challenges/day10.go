package challenges

import (
	"fmt"
	"log"
	"sort"

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
		x, _ := checkLine(line)
		score = score + x
	}
	return fmt.Sprintf("%d", score)
}

func (d *Day10) Part2() string {

	scores := make([]int, 0)
	for _, line := range d.rawInput {
		x, y := checkLine(line)
		if x != 0 {
			continue
		}
		scores = append(scores, y)

	}

	sort.Ints(scores)

	return fmt.Sprintf("%d", scores[len(scores)/2]) // WTF....

}

func checkLine(line string) (int, int) {

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
				return 3, 0
			}
		case ']':
			c, err := s.Pop()
			if err != nil {
				log.Panicf("stack errorr %c in %s", r, line)
			}
			if c != '[' {
				return 57, 0
			}
		case '}':
			c, err := s.Pop()
			if err != nil {
				log.Panicf("stack errorr %c in %s", r, line)
			}
			if c != '{' {
				return 1197, 0
			}
		case '>':
			c, err := s.Pop()
			if err != nil {
				log.Panicf("stack error %c in %s", r, line)
			}
			if c != '<' {
				return 25137, 0
			}
		default:
			log.Panicf("bad char %c in %s", r, line)
		}
	}

	score := 0
	for !s.IsEmpty() {
		c, _ := s.Pop()
		switch c {
		case '(':
			score = 5*score + 1
		case '[':
			score = 5*score + 2
		case '{':
			score = 5*score + 3
		case '<':
			score = 5*score + 4
		}
	}
	return 0, score
}
