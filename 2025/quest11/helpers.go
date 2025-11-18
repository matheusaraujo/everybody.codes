package main

import (
	"strconv"
)

func parseInput(puzzleInput []string) ([]int, int) {
	columns := make([]int, len(puzzleInput))
	sum := 0
	for i, a := range puzzleInput {
		columns[i], _ = strconv.Atoi(a)
		sum += columns[i]
	}

	return columns, sum / len(puzzleInput)
}
