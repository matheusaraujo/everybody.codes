package main

import (
	"strconv"
	"strings"
)

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func part3(puzzleInput []string) interface{} {
	names := strings.Split(puzzleInput[0], ",")
	instructions := strings.Split(puzzleInput[2], ",")
	position := 0

	for _, instruction := range instructions {
		direction := instruction[0]
		value, _ := strconv.Atoi(instruction[1:])

		switch direction {
		case 'L':
			position = abs(len(names) - value)
		case 'R':
			position = (value) % len(names)
		}
		names[0], names[position] = names[position], names[0]
	}

	return names[0]
}
