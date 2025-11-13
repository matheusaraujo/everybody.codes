package main

import (
	"strconv"
	"strings"
)

type line struct {
	a, b int
}

func parseInput(puzzleInput []string) []line {
	numbers := strings.Split(puzzleInput[0], ",")
	lines := []line{}

	for k := 1; k < len(numbers); k++ {
		a, _ := strconv.Atoi(numbers[k-1])
		b, _ := strconv.Atoi(numbers[k])
		lines = append(lines, line{a: min(a, b), b: max(a, b)})
	}
	return lines
}

func intersects(a, b, c, d int) bool {
	if a == c || a == d || b == c || b == d {
		return false
	}

	// two lines will intersect if exactly one point of the line is between the two points of the other line
	return isBetween(a, b, c) != isBetween(a, b, d)
}

func isBetween(a, b, x int) bool {
	return x > a && x < b
}
