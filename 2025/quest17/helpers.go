package main

import (
	"strconv"
)

func innerCircle(x, y, xc, yc, r int) bool {
	return (x-xc)*(x-xc)+(y-yc)*(y-yc) <= r*r
}

func atoi(line []string, x int, y int) int {
	a, _ := strconv.Atoi(string(line[x][y]))
	return a
}
