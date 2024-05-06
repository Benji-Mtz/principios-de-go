package main

import "fmt"

func main() {

	// Creating slices with MAKE -> slice_name := make([]type, length, capacity)
	// frutas := []string{"fresa", "limon"}
	frutas := make([]string, 0)
	fmt.Printf("Slice-Make: %v, len: %d, cap: %d\n", frutas, len(frutas), cap(frutas))

	frutas = append(frutas, "manzanas", "pera", "guayaba")
	fmt.Printf("Slice-Make-append: %v, len: %d, cap: %d\n\n", frutas, len(frutas), cap(frutas))

	// Slice: son apuntadores a un array, no poseen datos
	things := [7]string{"🚓", "🚚", "🚕", "🦄", "🚒", "🍅", "🎲"}
	cars := things[:5]
	last_4 := things[3:]
	all := things[:]
	last_4[1] = "🚑" // este valor se cambia en todos los arrays o slices por ser punteros a array

	fmt.Println("Things", things)
	fmt.Println("cars", cars)
	fmt.Println("last_4", last_4)

	fmt.Println("last_4[1]", last_4[1])
	fmt.Println("all", all)

	// Slices with len y cap
	animals := [5]string{"🐒", "🐕", "🦮", "🐦", "🐘"}
	pets := animals[1:3] // "🐕", "🦮"
	// Agregando elementos a un slice
	// pets = append(pets, "🐈", "🐕") // pets ahora es "🐕","🦮","🐈","🐕"-> y sustituira en animals "🐈" x "🐦" y "🐕" x "🐘" cuando no se supera el tamaño original

	// Cuando se supera el tamaño original se crea una nueva referencia y no se altera el array original
	// pets = append(pets, "🐈", "🐕", "🐱")
	// Array[4]{"🐕", "🦮",  "🐦", "🐘"} -> array len 2 y cap 4
	// Append
	// new Array[8]{"🐕", "🦮", "🐈", "🐕", "🐱"} array len 5 y  cap 8

	fmt.Println("\nanimals: ", animals)
	fmt.Println("pets: ", pets)
	fmt.Println("tamaño pets: ", len(pets))
	// La capacidad sera el resto de elementos que no hagarro el slice del array original
	fmt.Println("tamaño pets: ", cap(pets))

	// Otra forma de declarar slice mas directa es slice literal
	new_pets := []string{"🐕", "🐈"}
	fmt.Println("\nnew_pets", new_pets)
	fmt.Println("tamaño new_pets", len(new_pets))
	fmt.Println("capacidad new_pets", cap(new_pets))

	// Otra forma de declarar slice con make
	new_pets_make := make([]string, 0, 3)
	new_pets_make = append(new_pets_make, "🐈", "🐕", "🐱", "🐘")

	fmt.Println("\nnew_pets_make", new_pets_make)
	fmt.Println("tamaño new_pets_make", len(new_pets_make))
	fmt.Println("capacidad new_pets_make", cap(new_pets_make))

	// Valor cero de un slice
	var pets_nil []string //pets_nil := []string{} aqui daria falso el null

	fmt.Println("\npets_nil", pets_nil)
	fmt.Println("tamaño pets_nil", len(pets_nil))
	fmt.Println("capacidad pets_nil", cap(pets_nil))
	fmt.Println("Valor cero", pets_nil == nil)

}
