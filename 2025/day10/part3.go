package main

import (
	"fmt"
	"strings"
)

func part3(board []string) int {
	d := dragon(board)
	s := sheepColumns(board)
	m := map[string]int{}

	return dfs(board, m, d, s, 'S')
}

func sheepColumns(board []string) []int {
	s := make([]int, len(board[0]))
	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == 'S' {
			s[j] = 0
		} else {
			s[j] = -1
		}
	}
	return s
}

func encode(p position, s []int, t rune) string {
	buf := make([]string, len(s))
	for j := range s {
		buf[j] = fmt.Sprintf("%d", s[j])
	}
	return fmt.Sprintf("%d,%d|%s|%c", p.i, p.j, strings.Join(buf, ","), t)
}

func allEaten(s []int) bool {
	for _, v := range s {
		if v != -1 {
			return false
		}
	}
	return true
}

func copySheep(s []int) []int {
	w := make([]int, len(s))
	copy(w, s)
	return w
}

func dfs(board []string, m map[string]int, p position, s []int, t rune) int {
	if allEaten(s) {
		return 1
	}

	k := encode(p, s, t)
	if v, ok := m[k]; ok {
		return v
	}

	res := 0

	if t == 'S' {
		made := false

		for j, i := range s {
			if i == -1 {
				continue
			}

			ni := i + 1

			if ni >= len(board) {
				made = true
				continue
			}

			if p.i == ni && p.j == j && board[ni][j] != '#' {
				continue
			}

			made = true
			ns := copySheep(s)
			ns[j] = ni
			res += dfs(board, m, p, ns, 'D')
		}

		if !made {
			res += dfs(board, m, p, s, 'D')
		}

	} else {
		for _, mv := range moves {
			q := np(p, mv)

			if q.i < 0 || q.i >= len(board) || q.j < 0 || q.j >= len(board[0]) {
				continue
			}

			ns := copySheep(s)

			if ns[q.j] == q.i && board[q.i][q.j] != '#' {
				ns[q.j] = -1
			}

			res += dfs(board, m, q, ns, 'S')
		}
	}

	m[k] = res
	return res
}
