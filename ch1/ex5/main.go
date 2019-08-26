package main

import (
	"fmt"
	"math"
)

func powInt(num int, n int) int {
	return int(math.Pow(float64(num), float64(n)))
}

func main() {
	dice := []int{1, 2, 3, 4, 5, 6}
	ptns := []int{}
	for _, v := range dice {
		for _, w := range dice {
			ptns = append(ptns, v+w)
		}
	}
	cnt := 0
	for _, v := range ptns {
		if v == 6 || v == 8 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
