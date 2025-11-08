package main

func part1(puzzleInput []string) interface{} {
	teeth := parseInput(puzzleInput)
	var ratio float64 = 2025

	for i := 1; i < len(teeth); i++ {
		ratio *= float64(teeth[i]) / float64(teeth[i-1])
	}

	return int(ratio)
}
