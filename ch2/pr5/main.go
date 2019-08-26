package main

import "fmt"

func first() {
	a := 9 * 8 * 7 * 6 / (4 * 3 * 2)
	b := 5 * 4 * 3 / (3 * 2)
	c := 1
	res := a * b * c
	fmt.Println(res)
}

func second() {
	a := 9 * 8 * 7 / (3 * 2)
	b := 6 * 5 * 4 / (3 * 2)
	c := 1
	res := a * b * c
	fmt.Println(res)
}

func third() {
	a := 9 * 8 * 7 / (3 * 2)
	b := 6 * 5 * 4 / (3 * 2)
	c := 1
	res := a * b * c / (3 * 2)
	fmt.Println(res)
}

func forth() {
	a := 9 * 8 * 7 * 6 / (4 * 3 * 2)
	b := 4 * 3 / 2
	c := 1
	res := a * b * c / 2
	fmt.Println(res)
}

func main() {
	first()
	second()
	third()
	forth()
}
