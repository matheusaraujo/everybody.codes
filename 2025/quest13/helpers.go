package main

import (
	"strconv"
	"strings"
)

func transform(i, n int) int {
	if i%2 == 0 {
		return 1 + i/2
	}
	return n - i/2
}

func parseRange(s string) (int, int) {
	rng := strings.Split(s, "-")
	left, _ := strconv.Atoi(rng[0])
	right, _ := strconv.Atoi(rng[1])
	return left, right
}
