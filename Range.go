package main

import (
	"fmt"
)

func main() {
	var a = []int64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i, v := range a {
		fmt.Println("At index : ", i, "Value is : ", v)
	}
}
