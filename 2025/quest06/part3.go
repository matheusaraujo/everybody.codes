package main

func part3(puzzleInput []string) interface{} {
	line := puzzleInput[0]
	distance := 1000
	repeat := 1000
	arr := buildMap(line, distance, repeat)

	result := 0

	for i, c := range line {
		if c >= 'a' && c <= 'c' {
			result += arr[c-'a'][i]
		}
	}

	return result
}

func buildMap(line string, distance int, repeat int) [][]int {
	N := len(line)
	L, R := 0, repeat*N
	arr := make([][]int, 3)

	for i := 0; i < 3; i++ {
		arr[i] = make([]int, N)
	}

	for i, c := range line {
		if c >= 'A' && c <= 'C' {
			for k := 0; k < repeat; k++ {
				for j := 1; j <= distance; j++ {
					left := k*len(line) + (i - j)
					right := k*len(line) + (i + j)
					if left >= L {
						arr[c-'A'][left%N]++
					}
					if right < R {
						arr[c-'A'][right%N]++
					}
				}
			}
		}
	}

	return arr
}
