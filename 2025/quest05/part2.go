package main

import (
	"math"
)

func part2(puzzleInput []string) interface{} {
	mi, ma := math.MaxInt, math.MinInt
	for i := 0; i < len(puzzleInput); i++ {
		sword := buildSword(puzzleInput[i])
		mi = min(mi, sword.quality)
		ma = max(ma, sword.quality)
	}
	return ma - mi
}
