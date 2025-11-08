package main

import (
	"fmt"
	"math"
)

func part2(puzzleInput []string) interface{} {
	teeth := parseInput(puzzleInput)
	var ratio float64 = 1

	for i := 1; i < len(teeth); i++ {
		ratio *= float64(teeth[i]) / float64(teeth[i-1])
	}

	return fmt.Sprintf("%d", int(math.Ceil(10000000000000/ratio)))
}
