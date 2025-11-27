package main

func part2(puzzleInput []string) interface{} {
	plants, id, cases, _, _ := parseInput(puzzleInput)
	result := 0

	for _, c := range cases {
		for i := 0; i < len(c); i++ {
			p, _ := plants[i+1]
			p.thickness = c[i]
			plants[i+1] = p
		}
		e := plants[id].Energy(id, plants)
		if e > 0 {
			result += e
		}
	}

	return result
}
