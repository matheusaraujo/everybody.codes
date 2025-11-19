package main

func part1(puzzleInput []string) interface{} {
	barrels := parseInput(puzzleInput)
	return destroy(barrels, 0, 0)
}
