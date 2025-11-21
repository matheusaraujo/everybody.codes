package main

func parseInput(puzzleInput []string) [][]rune {
	out := make([][]rune, len(puzzleInput))
	for i, line := range puzzleInput {
		out[i] = []rune(line)
	}
	return out
}

func count(floor [][]rune) int {
	out := 0
	for i := 0; i < len(floor); i++ {
		for j := 0; j < len(floor[i]); j++ {
			if floor[i][j] == '#' {
				out++
			}
		}
	}
	return out
}

func hashFloor(floor [][]rune) string {
	out := make([]byte, 0, len(floor)*len(floor[0]))
	for i := range floor {
		for j := range floor[i] {
			out = append(out, byte(floor[i][j]))
		}
	}
	return string(out)
}

func next(floor [][]rune) [][]rune {
	nf := copyFloor(floor)
	for i := 0; i < len(floor); i++ {
		for j := 0; j < len(floor[i]); j++ {
			n := neighbors(floor, i, j)
			if floor[i][j] == '#' && n%2 == 0 {
				nf[i][j] = '.'
			}
			if floor[i][j] == '.' && n%2 == 0 {
				nf[i][j] = '#'
			}
		}
	}
	return nf
}

func neighbors(floor [][]rune, i, j int) int {
	out := 0
	deltas := [][2]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, d := range deltas {
		ni, nj := i+d[0], j+d[1]
		if ni >= 0 && nj >= 0 && ni < len(floor) && nj < len(floor[ni]) && floor[ni][nj] == '#' {
			out++
		}
	}
	return out
}

func copyFloor(floor [][]rune) [][]rune {
	out := make([][]rune, len(floor))
	for i := range floor {
		out[i] = make([]rune, len(floor[i]))
		copy(out[i], floor[i])
	}
	return out
}

func sum(s []int) int {
	out := 0
	for _, e := range s {
		out += e
	}
	return out
}
