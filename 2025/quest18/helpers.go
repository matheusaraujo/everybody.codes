package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Plant struct {
	thickness int
	branch    map[int]int
}

func parseInput(puzzleInput []string) (map[int]Plant, int, [][]int, []int, []int) {
	plants := make(map[int]Plant)
	cases := make([][]int, 0)
	positives := make([]int, 0)
	negatives := make([]int, 0)
	var id int

	plants[0] = Plant{thickness: 1, branch: make(map[int]int)}

	for _, line := range puzzleInput {
		if strings.HasPrefix(line, "Plant") {
			re := regexp.MustCompile(`Plant\s+(\d+)\s+with thickness\s+(-?\d+)`)
			matches := re.FindStringSubmatch(line)
			id, _ = strconv.Atoi(matches[1])
			thickness, _ := strconv.Atoi(matches[2])
			plants[id] = Plant{thickness: thickness, branch: make(map[int]int)}
		} else if strings.HasPrefix(line, "-") {
			re := regexp.MustCompile(`-\sfree\sbranch\swith\sthickness\s(-?\d+)`)
			matches := re.FindStringSubmatch(line)
			if len(matches) > 1 {
				thickness, _ := strconv.Atoi(matches[1])
				plants[id].branch[0] = thickness
			} else {
				re := regexp.MustCompile(`-\sbranch\sto\sPlant\s(\d+)\swith\sthickness\s(-?\d+)`)
				matches := re.FindStringSubmatch(line)
				id2, _ := strconv.Atoi(matches[1])
				thickness, _ := strconv.Atoi(matches[2])
				plants[id].branch[id2] = thickness
				if thickness > 0 {
					positives = append(positives, id2)
				} else {
					negatives = append(negatives, id2)
				}
			}
		} else if strings.HasPrefix(line, "0") || strings.HasPrefix(line, "1") {
			vs := strings.Split(line, " ")
			iv := make([]int, 0)
			for _, v := range vs {
				i, _ := strconv.Atoi(v)
				iv = append(iv, i)
			}
			cases = append(cases, iv)
		}
	}

	return plants, id, cases, positives, negatives
}

func (p Plant) Energy(id int, plants map[int]Plant) int {
	if len(p.branch) == 0 {
		return p.thickness
	}
	incomingEnergy := 0
	for k, v := range p.branch {
		incomingEnergy += plants[k].Energy(k, plants) * v
	}
	if incomingEnergy < p.thickness || p.thickness == 0 {
		return 0
	}
	return incomingEnergy
}
