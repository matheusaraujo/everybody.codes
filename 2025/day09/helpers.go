package main

import (
	"strconv"
	"strings"
)

type duck struct {
	id  int
	dna string
}

func newDuck(input string) duck {
	parts := strings.Split(input, ":")
	id, _ := strconv.Atoi(parts[0])
	dna := parts[1]
	return duck{id, dna}
}

func calc(c, p1, p2 string) int {
	s1, s2 := 0, 0
	for i := 0; i < len(c); i++ {
		if c[i] != p1[i] && c[i] != p2[i] {
			return 0
		}
		if c[i] == p1[i] {
			s1++
		}
		if c[i] == p2[i] {
			s2++
		}
	}
	return s1 * s2
}
