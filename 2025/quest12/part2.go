package main

func part2(puzzleInput []string) interface{} {
	barrels := parseInput(puzzleInput)
	return destroy(barrels, 0, 0) + destroy(barrels, len(barrels)-1, len(barrels[0])-1)
}
