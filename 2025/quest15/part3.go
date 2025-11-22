package main

import (
	"container/heap"
	"sort"
	"strconv"
	"strings"
)

func part3(puzzleInput []string) int {
	turns, amounts := parseInput3(puzzleInput)
	return solve(turns, amounts)
}

func parseInput3(puzzleInput []string) ([]byte, []int) {
	var turns []byte
	var amounts []int

	instructions := strings.Split(puzzleInput[0], ",")
	for _, inst := range instructions {
		turns = append(turns, inst[0])
		amount, _ := strconv.Atoi(string(inst[1:]))
		amounts = append(amounts, amount)
	}

	return turns, amounts
}

func solve(turns []byte, amounts []int) int {
	position := ORIGIN
	direction := UP

	var walls []Wall
	xs := NewSortedSet()
	ys := NewSortedSet()

	for i, turn := range turns {
		amount := amounts[i]

		if turn == 'R' {
			direction = direction.Clockwise()
		} else {
			direction = direction.CounterClockwise()
		}

		to := position.Add(direction.Scale(amount))

		if direction == LEFT || direction == RIGHT {
			x1, x2 := position.x, to.x
			minX, maxX := min(x1, x2), max(x1, x2)

			walls = append(walls, Wall{minX + 1, maxX - 1, position.y, position.y})

			xs.Insert(x1 - 1)
			xs.Insert(x2 + 1)
			ys.Insert(position.y - 1)
			ys.Insert(position.y + 1)
		} else if direction == UP || direction == DOWN {
			y1, y2 := position.y, to.y
			minY, maxY := min(y1, y2), max(y1, y2)

			walls = append(walls, Wall{position.x, position.x, minY + 1, maxY - 1})

			xs.Insert(position.x - 1)
			xs.Insert(position.x + 1)
			ys.Insert(y1 - 1)
			ys.Insert(y2 + 1)
		}

		position = to
	}

	xs.Insert(position.x)
	ys.Insert(position.y)

	collide := func(from, to Point) bool {
		x1, x2 := min(from.x, to.x), max(from.x, to.x)
		y1, y2 := min(from.y, to.y), max(from.y, to.y)

		for _, w := range walls {
			if !(x1 > w.bx || x2 < w.ax || y1 > w.by || y2 < w.ay) {
				return true
			}
		}
		return false
	}

	start := ORIGIN
	end := position

	todo := &MinHeap{}
	heap.Init(todo)

	seen := make(map[Point]bool)

	PushCost(todo, 0, start)
	seen[start] = true

	for todo.Len() > 0 {
		cost, from := PopCost(todo)

		if from == end {
			return cost
		}

		var nextPoints []Point

		if nextY, ok := ys.RangeNextBack(from.y); ok {
			nextPoints = append(nextPoints, Point{from.x, nextY})
		}

		if nextY, ok := ys.RangeNext(from.y); ok {
			nextPoints = append(nextPoints, Point{from.x, nextY})
		}

		if nextX, ok := xs.RangeNextBack(from.x); ok {
			nextPoints = append(nextPoints, Point{nextX, from.y})
		}

		if nextX, ok := xs.RangeNext(from.x); ok {
			nextPoints = append(nextPoints, Point{nextX, from.y})
		}

		for _, to := range nextPoints {
			_, exists := seen[to]
			if !collide(from, to) && !exists {
				seen[to] = true
				PushCost(todo, cost+from.Manhattan(to), to)
			}
		}
	}

	return -1
}

type SortedSet []int

func NewSortedSet(initial ...int) SortedSet {
	set := make(map[int]bool)
	for _, x := range initial {
		set[x] = true
	}
	var s SortedSet
	for x := range set {
		s = append(s, x)
	}
	sort.Ints(s)
	return s
}

func (s *SortedSet) Insert(x int) {
	i := sort.SearchInts(*s, x)
	if i < len(*s) && (*s)[i] == x {
		return
	}

	*s = append(*s, 0)
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = x
}

func (s SortedSet) RangeNext(start int) (int, bool) {
	i := sort.SearchInts(s, start+1)
	if i < len(s) {
		return s[i], true
	}
	return 0, false
}

func (s SortedSet) RangeNextBack(start int) (int, bool) {
	i := sort.SearchInts(s, start)

	if i < len(s) && s[i] == start {
		i--
	} else if i > 0 && i < len(s) {
		i--
	} else if i == len(s) && len(s) > 0 {
		i--
	}

	if i >= 0 {
		return s[i], true
	}
	return 0, false
}
