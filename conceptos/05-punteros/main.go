package main

import "fmt"

func main() {
	fmt.Print("Punteros - Initial values\n\n")
	fruit := "🍌"
	var puntero *string
	puntero = &fruit

	fmt.Printf("Valor - Tipo: %T, Valor: %s, Dirección: %v\n", fruit, fruit, &fruit)

	fmt.Printf("Puntero - Tipo: %T, Valor: %v, Desreferenciacion: %s\n\n",
		puntero,
		puntero,
		*puntero,
	)

	// sobreescribiendo el puntero
	*puntero = "🍍"

	fmt.Printf("Puntero - Tipo: %T, Valor: %v, Desreferenciacion: %s\n",
		puntero,
		puntero,
		*puntero,
	)

	fmt.Printf("Valor - Tipo: %T, Valor: %s, Dirección: %v\n\n", fruit, fruit, &fruit)

	a := 3

	fmt.Println("Antes de duplicar a=", a)
	Duplicar(&a)
	fmt.Println("Despues de duplicar a=", a)

}

func Duplicar(x *int) {
	*x *= 2
	fmt.Println("En la funcion duplicar x=", *x)
}
