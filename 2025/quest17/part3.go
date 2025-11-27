package main

import (
	"container/heap"
	"math"
)

type Cell struct {
	dist  int
	r     int
	c     int
	index int
}

// Min-Heap for Dijkstra
type PriorityQueue []*Cell

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Cell)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

var (
	R      int
	C      int
	G      [][]int
	vr, vc int
)

func dijkstra(startR, startC, rad int, goLeft bool) map[int]int {
	D := make(map[int]int)
	moves := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	pq := &PriorityQueue{}
	heap.Init(pq)

	heap.Push(pq, &Cell{dist: 0, r: startR, c: startC})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Cell)
		d, r, c := item.dist, item.r, item.c
		key := r*C + c

		if r < 0 || r >= R || c < 0 || c >= C {
			continue
		}

		if (r-vr)*(r-vr)+(c-vc)*(c-vc) <= rad*rad {
			continue
		}

		if r == vr {
			if goLeft && c > vc {
				continue
			}
			if !goLeft && c < vc {
				continue
			}
		}

		if _, ok := D[key]; ok {
			continue
		}

		D[key] = d

		for _, move := range moves {
			nr, nc := r+move[0], c+move[1]
			if nr < 0 || nr >= R || nc < 0 || nc >= C {
				continue
			}

			newDist := d + G[r][c]
			heap.Push(pq, &Cell{dist: newDist, r: nr, c: nc})
		}
	}
	return D
}

func part3(grid []string) interface{} {
	R = len(grid)
	C = len(grid[0])
	G = make([][]int, R)

	var sr, sc int

	for r, line := range grid {
		G[r] = make([]int, C)
		for c, char := range line {
			charStr := string(char)
			if charStr == "@" {
				vr, vc = r, c
			} else if charStr == "S" {
				sr, sc = r, c
			} else {
				G[r][c] = atoi(grid, r, c)
			}
		}
	}
	G[sr][sc] = 0

	for rad := 0; rad < 1000; rad++ {
		timeLimit := (rad+1)*30 - 1
		dLeft := dijkstra(sr, sc, rad, true)
		dRight := dijkstra(sr, sc, rad, false)

		score := math.MaxInt
		found := false

		for r := vr + 1; r < R; r++ {
			key := r*C + vc
			distLeft, okLeft := dLeft[key]
			distRight, okRight := dRight[key]

			if okLeft && okRight {
				currentScore := distLeft + distRight + G[r][vc]
				if currentScore < score {
					score = currentScore
					found = true
				}
			}
		}

		if found && score < timeLimit {
			return rad * score
		}
	}

	return 0
}
