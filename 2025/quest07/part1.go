package main

func part1(puzzleInput []string) interface{} {
	names, m := parseInput(puzzleInput)

	for _, name := range names {
		if isValid(name, m) {
			return name
		}
	}

	return nil
}
