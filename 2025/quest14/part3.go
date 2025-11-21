package main

func part3(puzzleInput []string) interface{} {
	pattern := parseInput(puzzleInput)
	floor := initial(34)

	limit := 1_000_000_000

	seen := make(map[string]int)
	contrib := []int{}

	for i := 0; i < limit; i++ {
		h := hashFloor(floor)

		// cycle detection
		if prev, ok := seen[h]; ok {
			return calc(prev, i, limit, contrib)
		}

		seen[h] = i
		floor = next(floor)
		if matches(floor, pattern) {
			contrib = append(contrib, count(floor))
		} else {
			contrib = append(contrib, 0)
		}
	}

	total := 0
	for _, v := range contrib {
		total += v
	}
	return total
}

func initial(n int) [][]rune {
	out := make([][]rune, n)
	for i := 0; i < n; i++ {
		out[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			out[i][j] = '.'
		}
	}
	return out
}

func matches(floor, pattern [][]rune) bool {
	shift := (len(floor) - len(pattern)) / 2
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[i]); j++ {
			if floor[shift+i][shift+j] != pattern[i][j] {
				return false
			}
		}
	}
	return true
}

func calc(prev, i, limit int, contrib []int) int {
	cycleStart := prev
	cycleLength := i - prev
	preCycleSum := sum(contrib[:cycleStart])
	cycleSum := sum(contrib[cycleStart:i])
	remaining := limit - cycleStart
	fullCycles := remaining / cycleLength
	tail := remaining % cycleLength
	total := preCycleSum + fullCycles*cycleSum + sum(contrib[cycleStart:cycleStart+tail])
	return total
}
