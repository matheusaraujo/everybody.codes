package main

import (
	"math"
)

func part2(puzzleInput []string) interface{} {
	sizes := parseInput(puzzleInput[0])
	mc := math.MaxInt

	for i := 0; i < len(sizes)-20; i++ {
		c := find(sizes, i)
		mc = min(mc, c)
	}

	return mc
}

func find(sizes []int, start int) int {
	sum, count := 0, 0
	if sizes[start+1] > sizes[start] {
		sum = sizes[start]
		count++
	}

	for i := start + 1; i < len(sizes); i++ {
		if sizes[i] > sizes[i-1] {
			sum += sizes[i]
			count++
			if count == 20 {
				return sum
			}
		}
	}

	return math.MaxInt
}
