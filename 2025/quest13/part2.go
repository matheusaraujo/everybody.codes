package main

func part2(puzzleInput []string) interface{} {
	turns := 20252025
	numbers := buildNumbers2(puzzleInput)
	spot := turns % len(numbers)
	return numbers[spot]
}

func buildNumbers2(puzzleInput []string) []int {
	left := make([]int, 1)
	right := make([]int, 0)
	left[0] = 1

	for i := 0; i < len(puzzleInput); i++ {
		r1, r2 := parseRange(puzzleInput[i])
		if i%2 == 0 {
			for i := r1; i <= r2; i++ {
				left = append(left, i)
			}
		} else {
			for i := r1; i <= r2; i++ {
				right = append([]int{i}, right...)
			}
		}
	}

	return append(left, right...)
}
