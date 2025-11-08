package main

func part1(puzzleInput []string) interface{} {
	sizes := parseInput(puzzleInput[0])

	sum := 0
	if sizes[1] > sizes[0] {
		sum = sizes[0]
	}

	for i := 1; i < len(sizes); i++ {
		if sizes[i] > sizes[i-1] {
			sum += sizes[i]
		}
	}

	return sum
}
