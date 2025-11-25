package main

func part3(puzzleInput []string) interface{} {
	blocks := parseInput(puzzleInput)
	spell := findSpell(blocks)

	n := 202520252025000
	return search(1, n/2, 0, 0, spell, n)
}

func search(minWall, maxWall, previousRes, previousWall int, spell []int, blocks int) int {
	if maxWall-minWall == 1 {
		result := count(maxWall, spell)
		if previousRes < result && result <= blocks {
			return maxWall
		}
		return previousWall
	}

	wall := (minWall + maxWall) / 2
	result := count(wall, spell)

	if result == blocks {
		return wall
	}
	if result > blocks {
		return search(minWall, wall, previousRes, previousWall, spell, blocks)
	}
	if result > previousRes {
		return search(wall, maxWall, result, wall, spell, blocks)
	}

	return search(wall, maxWall, previousRes, previousWall, spell, blocks)
}

func count(wall int, spell []int) int {
	s := 0
	for i := 0; i < len(spell); i++ {
		s += wall / spell[i]
	}
	return s
}
