package main

func part1(puzzleInput []string) interface{} {
	columns, _ := parseInput(puzzleInput)

	phase := 1
	for t := 0; t <= 10; t++ {
		moved := 0
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
		}
		if moved == 0 {
			phase = 2
		}
	}
	return checksum(columns)
}

func checksum(columns []int) int {
	result := 0
	for i := 0; i < len(columns); i++ {
		result += (i + 1) * columns[i]
	}
	return result
}
