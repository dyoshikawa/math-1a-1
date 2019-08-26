package main

import (
	"fmt"

	"github.com/dyoshikawa/math-1a-1/lib"
)

const (
	u    = 20
	aVal = 3
	bVal = 2
)

func main() {
	a := u / aVal
	lcmVal := lib.CalcLcm(aVal, bVal)
	aAndB := u / lcmVal
	onlyA := a - aAndB
	fmt.Println(u - onlyA)
}
