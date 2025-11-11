package main

func part2(puzzleInput []string) interface{} {
	A, B, C := 0, 0, 0
	result := 0
	for _, c := range puzzleInput[0] {
		switch c {
		case 'A':
			A++
		case 'a':
			result += A
		case 'B':
			B++
		case 'b':
			result += B
		case 'C':
			C++
		case 'c':
			result += C
		}
	}
	return result
}
