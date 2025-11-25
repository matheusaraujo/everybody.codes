package main

func part2(puzzleInput []string) interface{} {
	blocks := parseInput(puzzleInput)
	spell := findSpell(blocks)
	return prod(spell)
}
