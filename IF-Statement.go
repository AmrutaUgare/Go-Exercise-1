package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a number : ")
	var a int64
	fmt.Scanln(&a)

	if a > 0 {
		fmt.Println("Number is : ", a)
		a = a / 2
		fmt.Println("After number divided by 2 : ", a)
	} else {
		fmt.Println("Not greater than 0")
		fmt.Println("Number is : ", a)
	}
}
