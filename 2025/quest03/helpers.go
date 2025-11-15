package main

import (
	"slices"
	"strconv"
	"strings"
)

func parseInput(puzzleInput string) []int {
	strSizes := strings.Split(puzzleInput, ",")
	sizes := make([]int, len(strSizes))

	for i, s := range strSizes {
		n, _ := strconv.Atoi(s)
		sizes[i] = n
	}

	slices.Sort(sizes)

	return sizes
}
