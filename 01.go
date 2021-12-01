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

	_, err = lib.ParseLines2Int(data)
	if err != nil {
		log.Fatalf("parse error %+v", err)
	}

}
