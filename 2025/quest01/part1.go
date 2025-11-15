package main

import (
	"strconv"
	"strings"
)

func part1(puzzleInput []string) interface{} {
	names := strings.Split(puzzleInput[0], ",")
	instructions := strings.Split(puzzleInput[2], ",")
	position := 0

	for _, instruction := range instructions {
		direction := instruction[0]
		value, _ := strconv.Atoi(instruction[1:])

		switch direction {
		case 'L':
			position -= value
		case 'R':
			position += value
		}

		position = max(position, 0)
		position = min(position, len(names)-1)
	}

	return names[position]
}
