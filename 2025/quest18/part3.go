package main

func part3(puzzleInput []string) interface{} {
	plants, id, cases, positives, negatives := parseInput(puzzleInput)

	mxc := maxCase(positives, negatives, len(cases[0]))
	mx := simulate(plants, id, mxc)
	result := 0

	for _, c := range cases {
		for j := 0; j < len(c); j++ {
			p, _ := plants[j+1]
			p.thickness = c[j]
			plants[j+1] = p
		}
		e := plants[id].Energy(id, plants)
		if e > 0 {
			result += mx - e
		}
	}
	return result

}

// on the final input, each "root" plant contributes only positively or negatively to the next one
// so the maximum will be achieved when all positives are 1 and all negatives are 0
// this is not the case to the example, so this heuristic does not work for all cases
func maxCase(positives, negatives []int, ln int) []int {
	result := make([]int, ln)
	pp, pn := 0, 0
	for i := 1; i <= ln; i++ {
		if i == positives[pp] {
			result[i-1] = 1
			pp++
		} else if i == negatives[pn] {
			result[i-1] = 0
			pn++
		}
	}
	return result
}

func simulate(plants map[int]Plant, id int, c []int) int {
	for j := 0; j < len(c); j++ {
		p, _ := plants[j+1]
		p.thickness = c[j]
		plants[j+1] = p
	}
	return plants[id].Energy(id, plants)
}
