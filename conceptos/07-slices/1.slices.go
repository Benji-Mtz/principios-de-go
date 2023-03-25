package main

import "fmt"

func main() {
	// Slice: son apuntadores a un array, no poseen datos

	numbers := [7]string{"1", "2", "3", "4", "5", "6", "7"}
	first_5 := numbers[:5]
	last_4 := numbers[3:]
	all := numbers[:]
	last_4[1] = "cinco" // este valor se cambia en todos los arrays o slices

	fmt.Println("Numbers", numbers)
	fmt.Println("first_5", first_5)
	fmt.Println("last_4", last_4)

	fmt.Println("last_4[0]", last_4[0])

	fmt.Println("all", all)
}
