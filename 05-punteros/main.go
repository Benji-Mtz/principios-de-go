package main

import "fmt"

func main() {
	fmt.Println("Punteros")
	fruit := "🍌"
	var puntero *string
	puntero = &fruit

	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", fruit, fruit, &fruit)

	// sobreescribiendo el puntero
	*puntero = "🍍"

	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciacion: %s\n",
		puntero,
		puntero,
		*puntero,
	)

	a := 3

	fmt.Println("Antes de duplicar a=", a)
	Duplicar(&a)
	fmt.Println("Despues de duplicar a=", a)

}

func Duplicar(x *int) {
	*x *= 2
	fmt.Println("En la funcion duplicar x=", *x)
}
