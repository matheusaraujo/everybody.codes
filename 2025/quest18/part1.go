package main

func part1(puzzleInput []string) interface{} {
	plants, id, _, _, _ := parseInput(puzzleInput)
	return plants[id].Energy(id, plants)
}
