package main

import (
	"strconv"
	"strings"
)

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

type complexNumber struct {
	x, y int
}

func newComplexNumber(val string) complexNumber {
	p1 := strings.Index(val, "[")
	p2 := strings.Index(val, ",")
	p3 := strings.Index(val, "]")

	xStr := strings.TrimSpace(val[p1+1 : p2])
	yStr := strings.TrimSpace(val[p2+1 : p3])

	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)

	return complexNumber{x, y}
}

func sum(a, b complexNumber) complexNumber {
	return complexNumber{a.x + b.x, a.y + b.y}
}

func multiply(a, b complexNumber) complexNumber {
	return complexNumber{a.x*b.x - a.y*b.y, a.x*b.y + a.y*b.x}
}

func divide(a, b complexNumber) complexNumber {
	return complexNumber{a.x / b.x, a.y / b.y}
}

func shouldEngrave(a complexNumber) bool {
	r := newComplexNumber("[0,0]")
	d := newComplexNumber("[100000,100000]")

	for i := 0; i < 100; i++ {
		r = multiply(r, r)
		r = divide(r, d)
		r = sum(r, a)
		if max(abs(r.x), abs(r.y)) > 1000000 {
			return false
		}
	}

	return true
}
