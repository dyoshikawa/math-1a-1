package main

import "fmt"

// 4, 5, 5 a
// 5, 5, 6 b
// 4, 5, 6 c
func main() {
	a := 3
	b := 3
	c := 3 * 2
	odd := a + b
	even := c
	fmt.Println(odd)
	fmt.Println(even)
}
