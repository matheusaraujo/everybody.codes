package main

func part2(puzzleInput []string) interface{} {
	s := 0
	for c := 0; c < len(puzzleInput); c++ {
		for p1 := 0; p1 < len(puzzleInput); p1++ {
			for p2 := p1 + 1; p2 < len(puzzleInput); p2++ {
				if c != p1 && c != p2 {
					s += calc(newDuck(puzzleInput[c]).dna, newDuck(puzzleInput[p1]).dna, newDuck(puzzleInput[p2]).dna)
				}
			}
		}
	}
	return s
}
