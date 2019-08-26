package main

import (
	"fmt"

	"github.com/dyoshikawa/math-1a-1/lib"
)

const (
	u = 10
)

func first() {
	fmt.Println(lib.CombCnt(u, 4))
}

func second() {
	fmt.Println(lib.CombCnt(u, 5))
}

func third() {
	a := lib.CombCnt(u, 4)
	b := lib.CombCnt(u-4, 3)
	fmt.Println(a * b)
}

func main() {
	first()
	second()
	third()
}
