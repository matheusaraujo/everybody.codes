package main

func part2(puzzleInput []string) interface{} {
	a := newComplexNumber(puzzleInput[0])
	oa := sum(a, newComplexNumber("[1000,1000]"))

	result := 0

	for x := a.x; x <= oa.x; x += 10 {
		for y := a.y; y <= oa.y; y += 10 {
			if shouldEngrave(complexNumber{x, y}) {
				result += 1
			}
		}
	}

	return result
}
