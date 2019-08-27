package main

import (
	"fmt"
	"github.com/dyoshikawa/math-1a-1/lib"
)

func first() int {
	return lib.CombCnt(5, 2)
}

func second() []int {
	res1 := lib.Fact(3) * lib.CombCnt(4, 4) * lib.Fact(4)
	res2 := lib.Fact(4) * lib.Fact(4)
	return []int{res1, res2}
}

func main() {
	fmt.Println(first())
	fmt.Println(second())
}
