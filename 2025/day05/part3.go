package main

import (
	"slices"
)

func compare(a sword, b sword) int {
	if a.quality < b.quality {
		return 1
	} else if a.quality > b.quality {
		return -1
	} else {
		afb, bfb := a.fishbone, b.fishbone

		for {
			if afb == nil && bfb == nil {
				break
			}
			aa, bb := segmentNumber(afb), segmentNumber(bfb)

			if aa < bb {
				return 1
			} else if aa > bb {
				return -1
			} else {
				afb, bfb = afb.next, bfb.next
			}
		}

		if a.id > b.id {
			return -1
		}
		return 1
	}
}

func part3(puzzleInput []string) interface{} {
	swords := make([]sword, len(puzzleInput))
	for i := 0; i < len(puzzleInput); i++ {
		swords[i] = buildSword(puzzleInput[i])
	}

	slices.SortFunc(swords, compare)

	checksum := 0
	for i := 0; i < len(swords); i++ {
		checksum += (i + 1) * swords[i].id
	}

	return checksum
}
