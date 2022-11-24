package main

import (
	"fmt"
)

func main() {
	str := "Hello World"

	fmt.Println(str[0:5])
	fmt.Println(str[5:])

	a := str[0:5]
	b := str[5:]

	fmt.Println("String is : ", a, b)

}
