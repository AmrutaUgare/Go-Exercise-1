package main

import "fmt"

func main() {
	// defer runs at last of the program
	defer fmt.Println("Hello")
	defer fmt.Println("!")
	defer fmt.Println("World")

	/*defer fmt.Println("Bye.......................")
	fmt.Println("Hello")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")*/

}
