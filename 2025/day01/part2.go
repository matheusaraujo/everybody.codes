package main

import (
	"strconv"
	"strings"
)

func part2(puzzleInput []string) interface{} {
	names := strings.Split(puzzleInput[0], ",")
	instructions := strings.Split(puzzleInput[2], ",")
	position := 0

	for _, instruction := range instructions {
		direction := instruction[0]
		value, _ := strconv.Atoi(instruction[1:])

		switch direction {
		case 'L':
			position = (position - value) % len(names)
		case 'R':
			position = (position + value) % len(names)
		}
	}

	return names[position]
}
