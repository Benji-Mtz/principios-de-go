package main

import "fmt"

func main() {

	// frutas := []string{"fresa", "limon"}
	frutas := make([]string, 0)
	fmt.Println(frutas)

	frutas = append(frutas, "manzanas", "pera", "guayaba")
	fmt.Println(frutas)
}
