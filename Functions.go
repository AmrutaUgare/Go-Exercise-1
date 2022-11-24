package main

import (
	"fmt"
)

func main() {
	a := 50
	b := 20
	sum1(a, b)
	c := sub(a, b)
	fmt.Println("Substraction of ", a, " - ", b, " is ", c)
}

func sum1(a, b int) {

	fmt.Println("Sum of ", a, " + ", b, " is ", a+b)
	c := sum2((a + 2), (b + 2))
	fmt.Println("In Another way, sum of a nad b is : ", c)

}
func sum2(x, y int) int {
	return x + y

}

func sub(a, b int) int {
	return a - b
}
