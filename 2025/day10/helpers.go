package main

type position struct {
	i, j int
}

func dragon(board []string) position {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'D' {
				return position{i, j}
			}
		}
	}
	return position{-1, -1}
}

func isValid(p position, m int, n int) bool {
	return p.i >= 0 && p.i < m && p.j >= 0 && p.j < n
}

func np(p position, d position) position {
	return position{i: p.i + d.i, j: p.j + d.j}
}

var moves = []position{
	{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
	{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
}

func replaceAt(s string, i int, c rune) string {
	r := []rune(s)
	r[i] = c
	return string(r)
}
