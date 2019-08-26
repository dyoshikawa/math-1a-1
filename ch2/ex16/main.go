package main

import "fmt"

func first() {
	a := 6 * 5 / 2
	b := 4 * 3 / 2
	c := 2 / 2
	res := a * b * c
	fmt.Println(res)
}

func second() {
	a := 6 * 5 / 2
	b := 4 * 3 / 2
	c := 2 / 2
	res := a * b * c / (3 * 2)
	fmt.Println(res)
}

func main() {
	first()
	second()
}
