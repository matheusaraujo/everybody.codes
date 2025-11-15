package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func part3(puzzleInput []string) interface{} {
	teeth := parseInputPart3(puzzleInput)
	var ratio float64 = 100

	for i := 1; i < len(teeth); i++ {
		ratio *= float64(teeth[i-1][1]) / float64(teeth[i][0])
	}

	return fmt.Sprintf("%d", int(math.Floor(ratio)))
}

func parseInputPart3(puzzleInput []string) [][]int {
	teeth := make([][]int, len(puzzleInput))

	for i, s := range puzzleInput {
		ss := strings.Split(s, "|")

		if len(ss) == 1 {
			n, _ := strconv.Atoi(strings.TrimSpace(ss[0]))
			teeth[i] = []int{n, n}
		} else {
			n1, _ := strconv.Atoi(strings.TrimSpace(ss[0]))
			n2, _ := strconv.Atoi(strings.TrimSpace(ss[1]))
			teeth[i] = []int{min(n1, n2), max(n1, n2)}
		}
	}

	return teeth
}
