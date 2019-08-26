package main

import "fmt"

func first() {
	a := 6 * 5 * 4 / (3 * 2)
	b := 3 * 2 / 2
	c := 1
	res := a * b * c
	fmt.Println(res)
}

func second() {
	a := 6 * 5 * 4 * 3 / (4 * 3 * 2)
	b := 2
	c := 1
	res := a * b * c / 2
	fmt.Println(res)
}

func main() {
	first()
	second()
}
