package main

import "fmt"

func main() {
	// Puntero: variables que almacenan la dirección en memoria de un valor
	var color string = "🟥"
	// Declaramos el puntero
	var pointerColor *string
	// Asignamos la direccion de memoria de X variable
	pointerColor = &color

	// Modificando el varlo de la variable desde nuestro puntero
	*pointerColor = "ROJO"

	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", color, color, &color)                              // Tipo: string, Valor: 🟥, Dirección: 0xc00009e230
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", pointerColor, pointerColor, *pointerColor) // Tipo: *string, Valor: 0xc00009e230, Desreferenciación: 🟥
}
