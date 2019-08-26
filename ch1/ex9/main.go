package main

import (
	"fmt"

	"github.com/dyoshikawa/math-1a-1/lib"
)

func genAns(n int) int {
	divs := lib.Pf(n)
	m := make(map[int]int)
	for _, v := range divs {
		m[v]++
	}
	res := 1
	for _, v := range m {
		res *= (v + 1)
	}
	return res
}

func first() {
	a := 36
	fmt.Println(genAns(a))
}

func second() {
	a := 108
	fmt.Println(genAns(a))
}

func main() {
	first()
	second()
}
