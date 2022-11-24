package main

import "fmt"

func main() {
	students := []string{"Ram", "Shyam", "Ganesh", "Sachin", "Virat", "Sourabh"}

	display(students...)
}

func display(stud ...string) {
	for _, v := range stud {
		fmt.Println(v)
	}
}
