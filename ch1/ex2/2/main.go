package main

import (
	"fmt"

	"github.com/dyoshikawa/math-1a-1/lib"
)

func main() {
	u := 240
	a := 2
	b := 3
	aCnt := u / a
	lcmCnt := u / lib.CalcLcm(a, b)
	fmt.Println(aCnt - lcmCnt)
}
