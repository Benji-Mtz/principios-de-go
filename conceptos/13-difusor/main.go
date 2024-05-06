package main

import "fmt"

func main() {
	// Slices o porciones
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	concatenacion := []int{}

	concatenacion = append(s1, s2...)

	fmt.Print("Concatenacion de S1  y S2 => ", concatenacion)
}
