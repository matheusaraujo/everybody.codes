package main

type segment struct {
	left      int
	right     int
	clockwise bool
}

func (s segment) length() int {
	return s.right - s.left + 1
}

func part3(puzzleInput []string) interface{} {
	turns := 202520252025
	numbers, total := buildNumbers3(puzzleInput)
	slot, curr := turns%total, 0
	for _, seg := range numbers {
		d := slot - curr
		if d < seg.length() {
			if seg.clockwise {
				return seg.left + d
			}
			return seg.right - d
		}
		curr += seg.length()
	}
	return 0
}

func buildNumbers3(puzzleInput []string) ([]segment, int) {
	numbers := make([]segment, len(puzzleInput)+1)
	numbers[0] = s("1-1", true)
	total := 1

	for i := 0; i < len(puzzleInput); i++ {
		seg := s(puzzleInput[i], i%2 == 0)
		numbers[transform(i, len(puzzleInput))] = seg
		total += seg.length()
	}

	return numbers, total
}

func s(s string, clockwise bool) segment {
	left, right := parseRange(s)
	return segment{left, right, clockwise}
}
