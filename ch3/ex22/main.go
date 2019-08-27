package main

import (
	"fmt"
	"github.com/dyoshikawa/math-1a-1/lib"
)

func first() int {
	return lib.Fact(7) / (lib.Fact(2) * lib.Fact(2) * lib.Fact(3))
}

func second() int {
	return lib.Fact(6) / (lib.Fact(2) * lib.Fact(3))
}

func main() {
	fmt.Println(first())
	fmt.Println(second())
}
