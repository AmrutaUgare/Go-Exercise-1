package main

import (
	"fmt"
)

var a = 10

func main() {

	b := 22
	fmt.Println("Global variable , value of a : ", a)
	fmt.Println("Local variable , value of b : ", b)
	GetValue()
}

func GetValue() {
	c := 30
	// a is global variable so it can access all over program
	fmt.Println("Global variable , value of a : ", a)
	// not working bcoz b is a local variable in above function
	//fmt.Println("Local variable , value of b : ", b)
	fmt.Println("Local variable , value of c : ", c)

}
