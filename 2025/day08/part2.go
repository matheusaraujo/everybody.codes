package main

func part2(puzzleInput []string) interface{} {
	lines := parseInput(puzzleInput)
	intersections := 0

	for i := 1; i < len(lines); i++ {
		for j := 0; j < i; j++ {
			if intersects(lines[i].a, lines[i].b, lines[j].a, lines[j].b) {
				intersections++
			}
		}
	}

	return intersections
}
