package main

import (
	"fmt"

	"github.com/dyoshikawa/math-1a-1/lib"
)

func first() {
	a := 36
	divs := lib.Pf(a)
	m := make(map[int]int)
	for _, v := range divs {
		m[v]++
	}
	res := 1
	for k, v := range m {
		sum := 0
		sum += 1
		for i := 1; i <= v; i++ {
			sum += lib.PowInt(k, i)
		}
		res *= sum
	}
	fmt.Println(res)
}

func second() {
	a := 1500
	divs := lib.Pf(a)
	m := make(map[int]int)
	for _, v := range divs {
		m[v]++
	}

	resCnt := 1
	for _, v := range m {
		resCnt *= v + 1
	}
	fmt.Println(resCnt)

	resSum := 1
	for k, v := range m {
		sum := 0
		sum += 1
		for i := 1; i <= v; i++ {
			sum += lib.PowInt(k, i)
		}
		resSum *= sum
	}
	fmt.Println(resSum)
}

func main() {
	first()
	second()
}
