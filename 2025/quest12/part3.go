package main

func part3(puzzleInput []string) interface{} {
	barrels := parseInput(puzzleInput)
	result := 0
	for k := 0; k < 3; k++ {
		d, i, j := best(barrels)
		destroy(barrels, i, j)
		result += d
	}
	return result
}

func best(barrels [][]int) (int, int, int) {
	bd, bi, bj := -1, 0, 0

	for i := range barrels {
		for j := range barrels[i] {
			if barrels[i][j] == -1 {
				continue
			}

			tmp := copyBarrels(barrels)
			d := destroy(tmp, i, j)

			if d > bd {
				bd, bi, bj = d, i, j
			}
		}
	}
	return bd, bi, bj
}
