package main

import "fmt"

// P->Q rrrruuuu
func first() {
	res := 8 * 7 * 6 * 5 / (4 * 3 * 2)
	fmt.Println(res)
}

// P->A ruuu
// A->Q rrru
func second() {
	res := 4 * 4
	fmt.Println(res)
}

// P->B rrruuu
// A->Q ru
func third() {
	res := 6 * 5 * 4 / (3 * 2) * 2
	fmt.Println(res)
}

// P->A uuur
// A->B rr
// B->Q ru
func forth() {
	pabq := 4 * 1 * 2
	res := 16 + 40 - pabq
	fmt.Println(res)
}

func main() {
	first()
	second()
	third()
	forth()
}
