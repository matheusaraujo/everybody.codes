package main

func part3(puzzleInput []string) interface{} {
	prefixes, m := parseInput(puzzleInput)
	seen := make(map[string]struct{})

	for _, pref := range prefixes {
		if isValid(pref, m) {
			next([]rune(pref), m, seen)
		}
	}
	return len(seen)
}

func next(p []rune, m map[rune][]rune, seen map[string]struct{}) {
	n := m[p[len(p)-1]]

	if len(p) >= 7 && len(p) <= 11 {
		seen[string(p)] = struct{}{}
	}

	if len(p) < 11 {
		for _, q := range n {
			newP := append(append([]rune(nil), p...), q)
			next(newP, m, seen)
		}
	}
}
