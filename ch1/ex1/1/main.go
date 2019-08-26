package main

import (
	"fmt"

	"github.com/dyoshikawa/math-1a-1/lib"
)

func main() {
	u := 40
	a := 3
	b := 4
	aCnt := u / a
	bCnt := u / b
	lcm := lib.CalcLcm(a, b)
	dupCnt := u / lcm
	fmt.Println(aCnt + bCnt - dupCnt)
}
