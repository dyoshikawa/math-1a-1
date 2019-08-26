package main

import "fmt"

const (
	u = 100
	a = 80
	b = 45
	c = 40
)

func first() {
	aAndB := 34
	aOrB := a + b - aAndB
	fmt.Println(u - aOrB)
}

func second() {
	aAndNotC := 50
	aAndC := a - aAndNotC
	fmt.Println(c - aAndC)
}

func third() {
	notBAndNotC := 30
	bOrC := u - notBAndNotC
	bAndC := b + c - bOrC
	fmt.Println(bOrC - bAndC)
}

func main() {
	first()
	second()
	third()
}
