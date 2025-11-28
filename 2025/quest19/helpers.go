package main

import (
	"sort"
	"strconv"
	"strings"
)

type Interval struct{ Low, High int }
type GapData map[int][]Interval

func parseInput(puzzleInput []string) GapData {
	gaps := make(GapData)
	for _, line := range puzzleInput {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		gaps[x] = append(gaps[x], Interval{Low: y, High: y + z - 1})
	}
	return gaps
}

func solve(gaps GapData) int {
	x := 0

	yReachable := []Interval{{Low: 0, High: 0}}

	var xKeys []int
	for k := range gaps {
		xKeys = append(xKeys, k)
	}
	sort.Ints(xKeys)

	for _, xNext := range xKeys {
		var yReachableNext []Interval

		for _, yR := range yReachable {
			yLow := yR.Low
			yHigh := yR.High

			for _, yGap := range gaps[xNext] {
				yNextGapLow := yGap.Low
				yNextGapHigh := yGap.High

				yNextLow, yNextHigh := intersect(x, yLow, yHigh, xNext, yNextGapLow, yNextGapHigh)

				if yNextLow != -1 {
					if len(yReachableNext) > 0 && yNextLow <= yReachableNext[len(yReachableNext)-1].High+1 {
						lastIndex := len(yReachableNext) - 1
						yReachableNext[lastIndex].High = max(yReachableNext[lastIndex].High, yNextHigh)
					} else {
						yReachableNext = append(yReachableNext, Interval{Low: yNextLow, High: yNextHigh})
					}
				}
			}
		}

		x = xNext
		yReachable = yReachableNext

		if len(yReachable) == 0 {
			return -1
		}
	}

	y0Low := yReachable[0].Low
	return y0Low + (x-y0Low)/2
}

func intersect(x1, y1Low, y1High, x2, y2Low, y2High int) (int, int) {
	if x1 >= x2 {
		return -1, -1
	}

	dx := x2 - x1
	y3Low := y1Low - dx
	y3High := y1High + dx

	yLow := max(y2Low, y3Low)
	yLow += (x2 + yLow) & 1

	yHigh := min(y2High, y3High)
	yHigh -= (x2 + yHigh) & 1

	if yLow <= yHigh {
		return yLow, yHigh
	}

	return -1, -1
}
