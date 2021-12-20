package challenges

import (
	"fmt"

	"github.com/jcheyer/aoc2021/challenges/internal/imageenhancer"
	"github.com/jcheyer/aoc2021/lib"
)

type Day20 struct {
	rawInput []string
	algo     string
	rawImage []string
}

func (d *Day20) Load(file string) error {

	if len(d.rawInput) > 0 {
		return nil
	}
	var err error
	d.rawInput, err = lib.ReadInputLinesWithEmptyLines(file)
	if err != nil {
		return err
	}
	chunks := lib.Chunks(d.rawInput)
	d.algo = chunks[0][0]
	d.rawImage = chunks[1]

	return nil
}

func (d *Day20) Part1() string {
	i := imageenhancer.New(d.rawImage)
	i = i.Enhance(d.algo)
	i = i.Enhance(d.algo)
	return fmt.Sprintf("%d", i.Count())
}

func (d *Day20) Part2() string {
	img := imageenhancer.New(d.rawImage)
	for i := 0; i < 50; i++ {
		img = img.Enhance(d.algo)
	}
	return fmt.Sprintf("%d", img.Count())
}
