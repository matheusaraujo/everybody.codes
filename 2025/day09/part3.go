package main

func part3(puzzleInput []string) interface{} {
	f := make(map[int][]int)

	for c := 0; c < len(puzzleInput); c++ {
		for p1 := 0; p1 < len(puzzleInput); p1++ {
			for p2 := p1 + 1; p2 < len(puzzleInput); p2++ {
				if c != p1 && c != p2 {
					dc, dp1, dp2 := newDuck(puzzleInput[c]), newDuck(puzzleInput[p1]), newDuck(puzzleInput[p2])
					c := calc(dc.dna, dp1.dna, dp2.dna)
					if c > 0 {
						f[dc.id] = append(f[dc.id], dp1.id)
						f[dp1.id] = append(f[dp1.id], dc.id)

						f[dc.id] = append(f[dc.id], dp2.id)
						f[dp2.id] = append(f[dp2.id], dc.id)

						f[dp2.id] = append(f[dp2.id], dp1.id)
						f[dp1.id] = append(f[dp1.id], dp2.id)
					}
				}
			}
		}
	}
	return biggestFamily(f)
}

func biggestFamily(f map[int][]int) int {
	visited := make(map[int]bool)
	members, maxLen := 0, 0

	for node := range f {
		if !visited[node] {
			comp := []int{}
			stack := []int{node}
			visited[node] = true

			for len(stack) > 0 {
				cur := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				comp = append(comp, cur)

				for _, neigh := range f[cur] {
					if !visited[neigh] {
						visited[neigh] = true
						stack = append(stack, neigh)
					}
				}
			}

			if len(comp) > maxLen {
				maxLen = len(comp)
				members = 0
				for _, id := range comp {
					members += id
				}
			}
		}
	}

	return members
}
