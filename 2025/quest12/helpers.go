package main

import (
	"strconv"
)

func destroy(barrels [][]int, i, j int) int {
	if barrels[i][j] == -1 {
		return 0
	}

	b := barrels[i][j]
	barrels[i][j] = -1
	d := 1

	deltas := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, delta := range deltas {
		ni, nj := i+delta[0], j+delta[1]
		if ni < 0 || ni >= len(barrels) || nj < 0 || nj >= len(barrels[ni]) {
			continue
		}
		if barrels[ni][nj] != -1 && barrels[ni][nj] <= b {
			d += destroy(barrels, ni, nj)
		}
	}

	return d
}

func parseInput(puzzleInput []string) [][]int {
	barrels := make([][]int, len(puzzleInput))

	for i := 0; i < len(puzzleInput); i++ {
		barrels[i] = make([]int, len(puzzleInput[i]))
		for j := 0; j < len(puzzleInput[i]); j++ {
			barrels[i][j], _ = strconv.Atoi(string(puzzleInput[i][j]))
		}
	}

	return barrels
}

func copyBarrels(src [][]int) [][]int {
	dst := make([][]int, len(src))
	for i := range src {
		dst[i] = make([]int, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}
