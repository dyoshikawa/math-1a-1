package main

import "fmt"

func first() {
	res := 9 * 8 * 7 * 6 * 5 / (5 * 4 * 3 * 2 * 1)
	fmt.Println(res)
}

func second() {
	man := 5 * 4 * 3 / (3 * 2 * 1)
	woman := 4 * 3 / (2 * 1)
	res := man * woman
	fmt.Println(res)
}

func main() {
	first()
	second()
}
