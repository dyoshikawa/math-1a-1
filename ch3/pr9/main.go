package main

import (
	"fmt"
	"github.com/dyoshikawa/math-1a-1/lib"
)

func first() int {
	return lib.Fact(6) / lib.Fact(2)
}

func second() int {
	return lib.Fact(4) * lib.CombCnt(5, 2)
}

func main() {
	fmt.Println(first())
	fmt.Println(second())
}
