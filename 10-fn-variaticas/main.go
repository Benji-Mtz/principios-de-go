package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2))
}

func sum(nums ...int) int {
	total := 0

	for _, v := range nums {
		total += v
	}

	return total
}
