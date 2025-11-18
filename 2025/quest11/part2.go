package main

func part2(puzzleInput []string) interface{} {
	columns, _ := parseInput(puzzleInput)

	phase, t := 1, 0

	for t >= 0 {
		moved, first, balanced := 0, columns[0], 1
		for i := 1; i < len(columns); i++ {
			if phase == 1 && columns[i] < columns[i-1] {
				columns[i]++
				columns[i-1]--
				moved = 1
			} else if phase == 2 && columns[i] > columns[i-1] {
				columns[i]--
				columns[i-1]++
				moved = 1
			}
			if columns[i] != first {
				balanced = 0
			}
		}
		if balanced == 1 {
			return t - 1
		}
		if moved == 0 {
			phase = 2
		}
		t++
	}
	return 0
}
