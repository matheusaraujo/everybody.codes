package main

import (
	"math"
)

type Point struct {
	i, j int
}

type Queue []Point

func (q *Queue) Enqueue(p Point) {
	*q = append(*q, p)
}

func (q *Queue) Dequeue() Point {
	if len(*q) == 0 {
		return Point{-1, -1}
	}
	p := (*q)[0]
	*q = (*q)[1:]
	return p
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func part2(t []string) int {
	is, js := find(t, 'S')
	ie, je := find(t, 'E')

	start := Point{is, js}
	end := Point{ie, je}

	width := len(t[0])
	height := len(t)
	dist := make(map[int]int)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			dist[i*width+j] = math.MaxInt64
		}
	}

	queue := Queue{}

	dist[start.i*width+start.j] = 0
	queue.Enqueue(start)

	for !queue.IsEmpty() {
		current := queue.Dequeue()
		cDist := dist[current.i*width+current.j]

		if current == end {
			return cDist
		}

		moves := []Point{
			{0, -1},
			{0, 1},
		}

		deltaI := 1
		if (current.i+current.j)%2 == 0 {
			deltaI = -1
		}
		moves = append(moves, Point{deltaI, 0})

		for _, move := range moves {
			ni, nj := current.i+move.i, current.j+move.j

			if ni < 0 || nj < 0 || ni >= height || nj >= width {
				continue
			}

			nKey := ni*width + nj

			tile := t[ni][nj]
			if tile != 'T' && tile != 'E' {
				continue
			}

			if cDist+1 < dist[nKey] {
				dist[nKey] = cDist + 1
				queue.Enqueue(Point{ni, nj})
			}
		}
	}

	return -1
}

func find(t []string, g byte) (int, int) {
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			if t[i][j] == g {
				return i, j
			}
		}
	}
	return 0, 0
}
