package main

import (
	"fmt"
)

type house struct {
	noRooms int32
	price   float32

	city string
}

func main() {
	fmt.Println(house{noRooms: 2, price: 7000000.0, city: "Pune"})

	h := house{1, 4000000.0, "Pune"}
	fmt.Println(h)

	h.noRooms = 3
	h.price = 10000000.0
	h.city = "Pune"

	fmt.Println("Rooms : ", h.noRooms, "Price : ", h.price, "City : ", h.city)
}
