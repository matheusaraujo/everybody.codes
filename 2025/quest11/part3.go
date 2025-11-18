package main

// When it switches to phase 2, the columns are sorted.
// From that point on, we only need to count the distance to the equilibrium point
// (the sum of all positive numbers, or the sum of the absolute values ​​of the negative numbers)
// That's the number of rounds needed to reach equilibrium; each round moves one column towards equilibrium.
// The input for part 3 is already sorted, so we only need to count that number.
func part3(puzzleInput []string) interface{} {
	columns, avg := parseInput(puzzleInput)
	result := 0
	for i := 0; i < len(columns); i++ {
		result += max(0, avg-columns[i])
	}
	return result
}
