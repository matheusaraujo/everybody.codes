package main

import (
	"container/list"
)

func part3(puzzleInput []string) int {
	grid0 := toByteGrid(puzzleInput)
	grid1 := rotate(grid0, '.')
	grid2 := rotate(grid1, '.')
	grids := [][][]byte{grid0, grid1, grid2}

	h := len(grid0)
	w := len(grid0[0])

	seen := make([][]bool, 3)
	for i := 0; i < 3; i++ {
		seen[i] = make([]bool, h*w)
	}

	var sx, sy int
	found := false
	for y := 0; y < h && !found; y++ {
		for x := 0; x < w; x++ {
			if grid0[y][x] == 'S' {
				sx, sy = x, y
				found = true
				break
			}
		}
	}

	encode := func(x, y int) int { return y*w + x }

	q := list.New()
	q.PushBack([3]int{sx, sy, 0})
	seen[0][encode(sx, sy)] = true

	dirs := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
		{0, 0},
	}

	for q.Len() > 0 {
		cur := q.Remove(q.Front()).([3]int)
		x, y, cost := cur[0], cur[1], cur[2]
		layer := cost % 3

		if grids[layer][y][x] == 'E' {
			return cost
		}

		nextCost := cost + 1
		nextLayer := nextCost % 3

		for _, d := range dirs {
			nx, ny := x+d[0], y+d[1]

			if nx < 0 || ny < 0 || nx >= w || ny >= h {
				continue
			}

			if d[0] == 0 && d[1] == -1 {
				if !evenParity(x, y) {
					continue
				}
			}
			if d[0] == 0 && d[1] == 1 {
				if !oddParity(x, y) {
					continue
				}
			}

			cell := grids[nextLayer][ny][nx]
			if cell < 'A' || cell > 'Z' {
				continue
			}

			idx := encode(nx, ny)
			if seen[nextLayer][idx] {
				continue
			}

			seen[nextLayer][idx] = true
			q.PushBack([3]int{nx, ny, nextCost})
		}
	}

	return -1
}

func rotate(grid [][]byte, def byte) [][]byte {
	h := len(grid)
	w := len(grid[0])
	next := make([][]byte, h)
	for i := range next {
		next[i] = make([]byte, w)
		for j := range next[i] {
			next[i][j] = def
		}
	}

	for y := 0; y < h; y++ {
		idx := 0
		for x := y; x < w-y; x += 2 {
			toX := w - 1 - idx - 2*y
			toY := idx
			next[toY][toX] = grid[y][x]
			idx++
		}

		idx = 0
		for x := y + 1; x < w-y; x += 2 {
			toX := w - 1 - idx - 2*y - 1
			toY := idx
			next[toY][toX] = grid[y][x]
			idx++
		}
	}
	return next
}

func toByteGrid(lines []string) [][]byte {
	grid := make([][]byte, len(lines))
	for i := range lines {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func evenParity(x, y int) bool { return (x+y)%2 == 0 }
func oddParity(x, y int) bool  { return (x+y)%2 == 1 }
