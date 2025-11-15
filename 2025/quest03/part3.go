package main

func part3(puzzleInput []string) interface{} {
	sizes := parseInput(puzzleInput[0])

	curr, countCurr, maxCountCurr := sizes[0], 1, 1

	for i := 1; i < len(sizes); i++ {
		if sizes[i] == curr {
			countCurr++
		} else {
			curr = sizes[i]
			countCurr = 1
		}
		maxCountCurr = max(maxCountCurr, countCurr)
	}

	return maxCountCurr
}
