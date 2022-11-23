package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter your name : ")

	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello", name)

	fmt.Println("Enter a number : ")
	var a int
	fmt.Scanln(&a)

	if a >= 1 && a <= 10 {
		fmt.Println(a, " is in between 1 to 10 ")
	} else {
		fmt.Println(a, " is not in between 1 to 10")
	}
}
