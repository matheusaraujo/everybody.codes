package main

import (
	"fmt"
)

func part1(puzzleInput []string) interface{} {
	a := newComplexNumber(puzzleInput[0])
	r := newComplexNumber("[0,0]")
	d := newComplexNumber("[10,10]")

	for i := 0; i < 3; i++ {
		r = multiply(r, r)
		r = divide(r, d)
		r = sum(r, a)
	}
	return fmt.Sprintf("[%d,%d]\n", r.x, r.y)
}
