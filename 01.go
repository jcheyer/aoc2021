package main

import (
	"log"

	"github.com/jcheyer/aoc2021/lib"
)

func day01_1(file string) {
	data, err := lib.ReadInputLines(file)
	if err != nil {
		log.Fatalf("input error %+v", err)
	}

	sonarData, err := lib.ParseLines2Int(data)
	if err != nil {
		log.Fatalf("parse error %+v", err)
	}

	var prev int64 = sonarData[0]
	var counter int64 = 0
	for i := 1; i < len(sonarData); i++ {
		if sonarData[i] > prev {
			counter++
		}
		prev = sonarData[i]
	}
	log.Printf("Day 1 Part 1 Result: %d", counter)

}

func day01_2(file string) {
	data, err := lib.ReadInputLines(file)
	if err != nil {
		log.Fatalf("input error %+v", err)
	}

	sonarData, err := lib.ParseLines2Int(data)
	if err != nil {
		log.Fatalf("parse error %+v", err)
	}

	var prev int64 = sonarData[0] + sonarData[1] + sonarData[2]
	var counter int64 = 0
	for i := 1; i < len(sonarData)-2; i++ {
		window := sonarData[i] + sonarData[i+1] + sonarData[i+2]
		if window > prev {
			counter++
		}
		prev = window
	}

	log.Printf("Day 1 Part 2 Result: %d", counter)

}
