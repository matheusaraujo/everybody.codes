package main

import (
	"strings"
)

func parseInput(puzzleInput []string) ([]string, map[rune][]rune) {
	names := strings.Split(puzzleInput[0], ",")
	m := buildMap(puzzleInput[2:])
	return names, m
}

func buildMap(rules []string) map[rune][]rune {
	m := make(map[rune][]rune)
	for _, rule := range rules {
		parts := strings.Split(rule, " > ")
		left := rune(parts[0][0])
		right := []rune(strings.ReplaceAll(strings.TrimSpace(parts[1]), ",", ""))
		m[left] = append(m[left], right...)
	}
	return m
}

func isValid(name string, rules map[rune][]rune) bool {
	runes := []rune(name)
	for i := 1; i < len(runes); i++ {
		prev := runes[i-1]
		curr := runes[i]

		allowed  := rules[prev]

		valid := false
		for _, r := range allowed {
			if r == curr {
				valid = true
				break
			}
		}

		if !valid {
			return false
		}
	}
	return true
}