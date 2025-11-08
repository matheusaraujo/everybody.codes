package main

import (
	"strconv"
	"strings"
)

type segment struct {
	left, center, right int
	next                *segment
}

type sword struct {
	id       int
	quality  int
	fishbone *segment
}

func segmentNumber(seg *segment) int {
	str := ""

	if seg.left != 0 {
		str += strconv.Itoa(seg.left)
	}
	if seg.center != 0 {
		str += strconv.Itoa(seg.center)
	}
	if seg.right != 0 {
		str += strconv.Itoa(seg.right)
	}

	if str == "" {
		return 0
	}

	num, _ := strconv.Atoi(str)
	return num
}

func buildFishbone(s *segment, n int) *segment {
	if s == nil {
		s = &segment{}
	}

	if s.center == 0 {
		s.center = n
	} else if s.left == 0 && n < s.center {
		s.left = n
	} else if s.right == 0 && n > s.center {
		s.right = n
	} else {
		s.next = buildFishbone(s.next, n)
	}

	return s
}

func buildSword(input string) sword {
	line := strings.Split(input, ":")
	id, numbers := line[0], strings.Split(line[1], ",")

	var root *segment

	for i := 0; i < len(numbers); i++ {
		n, _ := strconv.Atoi(numbers[i])
		root = buildFishbone(root, n)
	}

	head := root

	result := ""

	for {
		if root == nil {
			break
		}
		s := strconv.Itoa(root.center)
		result += s
		root = root.next
	}

	quality, _ := strconv.Atoi(result)
	_id, _ := strconv.Atoi(id)
	return sword{_id, quality, head}
}
