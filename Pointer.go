package main

import (
	"fmt"
)

func main() {
	a := 10
	fmt.Println("Value of a : ", a)
	fmt.Println("Address of a : ", &a)

	var b *int

	b = &a
	fmt.Println("Address of b : ", b)

	c := *b
	fmt.Println("Value of c : ", c)
}
