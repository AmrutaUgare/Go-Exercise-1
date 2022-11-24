package main

import "fmt"

func main() {
	m1 := map[string]int{}

	m1["Pune"] = 1
	m1["Mumbai"] = 2
	m1["Kolhapur"] = 3
	m1["Delhi"] = 4
	m1["Jaipur"] = 5

	fmt.Println(m1)

	// display value of perticula key
	fmt.Println("Value of m1[Delhi] is : ", m1["Delhi"])

	// add in map
	m1["Rajastan"] = 6
	m1["Banglore"] = 7

	fmt.Println(m1)

	// show map using range

	for i, v := range m1 {
		fmt.Println("Index : ", i, "||", "Value : ", v)
	}

	// delete

	delete(m1, "Mumbai")
	fmt.Println("After deletion ", m1)

}
