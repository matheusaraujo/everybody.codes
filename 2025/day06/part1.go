package main

func part1(puzzleInput []string) interface{} {
	A := 0
	result := 0
	for _, c := range puzzleInput[0] {
		switch c {
		case 'A':
			A++
		case 'a':
			result += A
		}
	}
	return result
}
