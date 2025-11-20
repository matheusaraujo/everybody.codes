package main

import (
	"strconv"
)

func part1(puzzleInput []string) interface{} {
	turns := 2025
	numbers := buildNumbers1(puzzleInput)
	spot := turns % len(numbers)
	return numbers[spot]
}

func buildNumbers1(puzzleInput []string) []int {
	n := len(puzzleInput)
	numbers := make([]int, n+1)
	numbers[0] = 1

	for i := 0; i < n; i++ {
		numbers[transform(i, n)], _ = strconv.Atoi(puzzleInput[i])
	}

	return numbers
}
