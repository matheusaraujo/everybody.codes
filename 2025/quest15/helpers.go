package main

import (
	"container/heap"
	"strconv"
	"strings"
)

// --- Structs
type Point struct {
	x, y int
}

type Wall struct {
	ax, bx, ay, by int
}

func (a Point) Equals(b Point) bool {
	return a.x == b.x && a.y == b.y
}

func (p Point) Add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func (p Point) Scale(s int) Point {
	return Point{p.x * s, p.y * s}
}

func (p Point) Manhattan(other Point) int {
	return abs(p.x-other.x) + abs(p.y-other.y)
}

func (d Point) Clockwise() Point {
	// (x, y) -> (y, -x)
	return Point{d.y, -d.x}
}

func (d Point) CounterClockwise() Point {
	// (x, y) -> (-y, x)
	return Point{-d.y, d.x}
}

var (
	ORIGIN = Point{0, 0}
	UP     = Point{0, 1}
	DOWN   = Point{0, -1}
	LEFT   = Point{-1, 0}
	RIGHT  = Point{1, 0}
)

// --- math
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// --- buildMap (part1 and part2)
func buildMap(puzzleInput []string) (map[Point]bool, Point, int, int, int, int) {
	m := make(map[Point]bool)
	instructions := strings.Split(puzzleInput[0], ",")
	x, y, dir := 0, 0, 0
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for _, inst := range instructions {
		d := inst[0]
		length, _ := strconv.Atoi(string(inst[1:]))
		if d == 'L' {
			dir = dir - 1
			if dir == -1 {
				dir = 3
			}
		} else if d == 'R' {
			dir = (dir + 1) % 4
		}
		for l := 0; l < length; l++ {
			if dir == 0 {
				x--
			} else if dir == 1 {
				y++
			} else if dir == 2 {
				x++
			} else if dir == 3 {
				y--
			}
			minX = min(minX, x)
			maxX = max(maxX, x)
			minY = min(minY, y)
			maxY = max(maxY, y)
			m[Point{x, y}] = true
		}
	}
	return m, Point{x, y}, minX, maxX, minY, maxY
}

// --- bfs (part1 and part2)
func bfs(m map[Point]bool, start, end Point, minX, maxX, minY, maxY int) int {
	dist := map[Point]int{start: 0}
	queue := []Point{start}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, d := range []Point{UP, DOWN, LEFT, RIGHT} {
			n := cur.Add(d)

			if n.Equals(end) {
				return dist[cur] + 1
			}

			if m[n] {
				continue
			}

			if _, ok := dist[n]; ok {
				continue
			}

			if n.x < minX || n.x > maxX || n.y < minY || n.y > maxY {
				continue
			}

			dist[n] = dist[cur] + 1
			queue = append(queue, n)
		}
	}

	return -1
}

// --- MinHeap Implementation (Dijkstra on part3)
type Item struct {
	cost  int
	point Point
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func PushCost(h *MinHeap, cost int, p Point) {
	heap.Push(h, Item{cost, p})
}

func PopCost(h *MinHeap) (int, Point) {
	if h.Len() == 0 {
		return -1, Point{}
	}
	item := heap.Pop(h).(Item)
	return item.cost, item.point
}
