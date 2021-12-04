package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jcheyer/aoc2021/challenges"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide the problem name (e.g. 'day01') as parameter.")
	}

	challengeName := os.Args[1]

	challengeMap := make(map[string]challenges.Challenge)

	challengeMap["day01"] = &challenges.Day01{}
	challengeMap["day02"] = &challenges.Day02{}
	challengeMap["day03"] = &challenges.Day03{}
	challengeMap["day04"] = &challenges.Day04{}
	challengeMap["day05"] = &challenges.Day04{}

	problem, defined := challengeMap[challengeName]
	if !defined {
		log.Panicf("challenge not found: %s\n", challengeName)
	}

	if err := problem.Load("input/" + challengeName + ".txt"); err != nil {
		log.Panicf("load error(%s): %+v\n", challengeName, err)
	}

	fmt.Printf("Part1: %s\n", problem.Part1())
	fmt.Printf("Part2: %s\n", problem.Part2())

}
