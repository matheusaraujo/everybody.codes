package main

func part1(puzzleInput []string) interface{} {
	output, columns := 0, 90
	blocks := parseInput(puzzleInput)
	for _, b := range blocks {
		output += columns / b
	}
	return output
}
