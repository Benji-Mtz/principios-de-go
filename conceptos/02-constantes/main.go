package main

import "fmt"

const (
	os     = "linux"
	domain = "benji.com"
)

// Ejemplo con iota tipo ENUM donde incrementa de 1 en 1 aunque inicia en 0
// Constantes a diferencia de las variables se pueden usar o no
const (
	Jan = iota + 1
	Feb
	Mar
	Abr
	May
	Jun
)

func main() {
	const pi = 3.1416
	// pi = 4.15 //esto mandaria un error
	fmt.Println(pi)
	fmt.Println(os, domain)
	fmt.Println(
		Jan,
		Feb,
		Mar,
		Abr,
		May,
		Jun,
	)
}
