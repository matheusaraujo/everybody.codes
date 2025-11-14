package main

func part1(puzzleInput []string) interface{} {
	d1, d2, d3 := newDuck(puzzleInput[0]), newDuck(puzzleInput[1]), newDuck(puzzleInput[2])
	return max(calc(d1.dna, d2.dna, d3.dna), calc(d2.dna, d1.dna, d3.dna), calc(d3.dna, d1.dna, d2.dna))
}
