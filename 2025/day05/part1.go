package main

func part1(puzzleInput []string) interface{} {
	sword := buildSword(puzzleInput[0])
	return sword.quality
}
