package main

func part3(puzzleInput []string) interface{} {
	lines := parseInput(puzzleInput)
	maxIntersections, n := 0, 256

	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			intersections := 0
			for k := 0; k < len(lines); k++ {
				if intersects(i, j, lines[k].a, lines[k].b) || (i == lines[k].a && j == lines[k].b) {
					intersections++
				}
			}
			maxIntersections = max(maxIntersections, intersections)
		}
	}

	return maxIntersections
}
