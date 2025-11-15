package main

func part2(board []string) interface{} {
	turns := 20

	d := dragonPositions(turns, len(board), len(board[0]))
	eaten := make(map[int]bool)

	for idx, s := range sheep(board) {
		for t := 1; t <= turns; t++ {
			i1 := s.i + (t - 1)
			i2 := s.i + t
			if eaten[idx] || i1 >= len(board) || i2 >= len(board) {
				break
			}

			if (d[t][position{i1, s.j}] && board[i1][s.j] != '#') ||
				(d[t][position{i2, s.j}] && board[i2][s.j] != '#') {
				eaten[idx] = true
			}
		}
	}

	return len(eaten)
}

func dragonPositions(turns, m, n int) []map[position]bool {
	d := make([]map[position]bool, turns+1)
	d[0] = map[position]bool{{i: m / 2, j: n / 2}: true}
	for t := 1; t <= turns; t++ {
		d[t] = make(map[position]bool)
		for pos := range d[t-1] {
			for _, mv := range moves {
				nPos := np(pos, mv)
				if isValid(nPos, m, n) {
					d[t][nPos] = true
				}
			}
		}
	}
	return d
}

func sheep(board []string) []position {
	s := make([]position, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'S' {
				s = append(s, position{i, j})
			}
		}
	}
	return s
}
