package main

func part1(t []string) interface{} {
	s := 0
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(t[i]); j++ {
			if t[i][j] == 'T' {
				if j > 0 && t[i][j-1] == 'T' {
					s++
				}
				if j < len(t[i])-1 && t[i][j+1] == 'T' {
					s++
				}
				if (i+j)%2 == 0 && i > 0 && t[i-1][j] == 'T' {
					s++
				}
				if (i+j)%2 == 1 && i < len(t) && t[i+1][j] == 'T' {
					s++
				}
			}
		}
	}
	return s / 2
}
