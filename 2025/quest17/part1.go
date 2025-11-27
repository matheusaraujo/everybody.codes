package main

func part1(grid []string) interface{} {
	n := len(grid)
	xc, yc, res := n/2, n/2, 0

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if innerCircle(x, y, xc, yc, 10) {
				res += atoi(grid, x, y)
			}
		}
	}

	return res
}
