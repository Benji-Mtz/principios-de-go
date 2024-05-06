package main

import "fmt"

func main() {
	// var food [3]string
	// food[0], food[1], food[2] = "🍔", "🍕", "🌭"

	// var food = [3]string{"🍔", "🍕", "🌭"}

	// [...] es solo para no estar alterando manualmente el tamaño del array
	food := [...]string{"🍔", "🍕", "🌭", "🍉"}
	fmt.Println(food)

}
