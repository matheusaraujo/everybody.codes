package main

import (
	"strconv"
	"strings"
)

func part1(puzzleInput []string) interface{} {
	numbers := strings.Split(puzzleInput[0], ",")
	times := 0
	for k := 1; k < len(numbers); k++ {
		i, _ := strconv.Atoi(numbers[k-1])
		j, _ := strconv.Atoi(numbers[k])
		if passThroughTheCenter(i, j, 32) {
			times++
		}
	}
	return times
}

func passThroughTheCenter(i, j, n int) bool {
	return abs(i-j) == n/2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
