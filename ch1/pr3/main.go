package main

import "fmt"

var (
	dice = []int{1, 2, 3, 4, 5, 6}
)

func first() {
	cnt := 0
	for _, v := range dice {
		for _, w := range dice {
			if v+w == 5 || v+w == 6 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}

func second() {
	evenCnt := 0
	for _, v := range dice {
		if v%2 == 0 {
			evenCnt++
		}
	}
	fmt.Println(evenCnt * evenCnt)
}

func main() {
	first()
	second()
}
