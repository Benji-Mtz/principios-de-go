package main

import "fmt"

func main() {
	/*
		var (
			apple      string = "🍎"
			banana     string = "🍌"
			strawberry string = "🍓"
		)
	*/

	// var apple, banana, strawberry string = "🍎", "🍌", "🍓"
	// var apple, banana, strawberry = "🍎", "🍌", "🍓"
	apple, banana, strawberry := "🍎", "🍌", "🍓"
	apple, lemon := "manzana", "🍋"

	fmt.Println(apple, banana, strawberry, lemon)
}
