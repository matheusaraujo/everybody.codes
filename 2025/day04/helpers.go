package main

import (
	"slices"
	"strconv"
)

func parseInput(puzzleInput []string) []int {
	teeth := make([]int, len(puzzleInput))

	for i, s := range puzzleInput {
		n, _ := strconv.Atoi(s)
		teeth[i] = n
	}

	slices.Sort(teeth)

	return teeth
}
