package main

import (
	"fmt"
)

func main() {
	strs := []string{"a", "b", "c", "d"}
	res := [][]string{}
	for _, v := range strs {
		for _, w := range strs {
			for _, x := range strs {
				if v == w || v == x || w == x {
					continue
				}
				res = append(res, []string{v, w, x})
			}
		}
	}
	fmt.Println(len(res))
}
