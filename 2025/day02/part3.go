package main

func part3(puzzleInput []string) interface{} {
	a := newComplexNumber(puzzleInput[0])
	oa := sum(a, newComplexNumber("[1000,1000]"))

	result := 0

	for x := a.x; x <= oa.x; x += 1 {
		for y := a.y; y <= oa.y; y += 1 {
			if shouldEngrave(complexNumber{x, y}) {
				result += 1
			}
		}
	}

	return result
}
