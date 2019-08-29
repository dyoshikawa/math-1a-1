package main

import (
	"fmt"
)

func first() int {
	ptns := [][]int{}
	for i := 1; i <= 9; i++ {
		for j := 1; j < 9-i; j++ {
			all := 9
			ptn := []int{}
			for k := 0; k < 3; k++ {
				switch k {
				case 0:
					all -= i
					ptn = append(ptn, i)
				case 1:
					all -= j
					ptn = append(ptn, j)
				case 2:
					ptn = append(ptn, all)
				}
			}
			if ptn[1] < ptn[0] || ptn[2] < ptn[1] {
				continue
			}
			flg := false
			for _, v := range ptn {
				if v <= 0 || 4 < v {
					flg = true
					break
				}
			}
			if flg {
				continue
			}
			ptns = append(ptns, ptn)
		}
	}
	return len(ptns)
}

func second() {

}

func main() {
	fmt.Println(first())
}
