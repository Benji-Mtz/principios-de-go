package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello ", name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Println("Your age is ", age)
}
