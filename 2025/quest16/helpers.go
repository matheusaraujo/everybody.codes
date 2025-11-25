package main

import (
	"strconv"
	"strings"
)

func parseInput(puzzleInput []string) []int {
	output := make([]int, 0)
	line := strings.Split(puzzleInput[0], ",")

	for i := 0; i < len(line); i++ {
		n, _ := strconv.Atoi(line[i])
		output = append(output, n)
	}

	return output
}

func prod(arr []int) int {
	s := 1
	for _, n := range arr {
		s *= n
	}
	return s
}

func findSpell(blocks []int) []int {
	output := make([]int, 0)
	blocks = append([]int{0}, blocks...)

	for k := 1; k < len(blocks); k++ {
		if blocks[k] > 0 {
			output = append(output, k)

			for j := 1; k*j < len(blocks); j++ {
				blocks[j*k]--
			}
		}
	}

	return output
}
