package main

func part2(puzzleInput []string) interface{} {
	names, m := parseInput(puzzleInput)
	result := 0
	for i, name := range names {
		if isValid(name, m) {
			result += (i + 1)
		}
	}

	return result
}
