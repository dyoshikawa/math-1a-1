package main

import (
	"fmt"
	"github.com/dyoshikawa/math-1a-1/lib"
	"reflect"
)

func first() int {
	a := lib.Fact(4) / lib.Fact(2) / lib.Fact(2)
	b := lib.CombCnt(5, 3)
	return a * b
}

func second() int {
	strs := []string{"A", "A", "B", "B"}
	ptns := [][]string{}
	for k, v := range strs {
		for l, w := range strs {
			for m, x := range strs {
				for n, y := range strs {
					if k == l || k == m || k == n || l == m || l == n || m == n {
						continue
					}
					newPtn := []string{v, w, x, y}
					flg := false
					for _, z := range ptns {
						if reflect.DeepEqual(newPtn, z) {
							flg = true
							break
						}
					}
					if flg {
						continue
					}
					ptns = append(ptns, newPtn)
				}
			}
		}
	}
	cnt := 0
	for _, v := range ptns {
		c := 3
		conti := 0
		prev := ""
		for _, w := range v {
			if prev == w {
				c--
				conti++
			}
			prev = w
		}
		free := 5 - conti
		cnt += lib.CombCnt(free, c)
	}
	return cnt
}

func main() {
	fmt.Println(first())
	fmt.Println(second())
}
