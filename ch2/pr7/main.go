package main

import "fmt"

// P->Q rrrrruuuu
func first() int {
	res := 9 * 8 * 7 * 6 / (4 * 3 * 2)
	return res
}

// P->R rruuu
// R->Q rrru
func second() int {
	r := 5 * 4 / 2 * 4
	return first() - r
}

// P->S1 rrruu
// S1->S2 r
// S2->Q ruu
func third() int {
	ps1 := 5 * 4 / 2
	s1s2 := 1
	s2q := 3
	return ps1 * s1s2 * s2q
}

func main() {
	fmt.Println(first())
	fmt.Println(second())
	fmt.Println(third())
}
