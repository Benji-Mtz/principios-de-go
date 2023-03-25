package main

import "fmt"

func main() {

	//`byte` alias for uint8
	//`rune` alias for int32

	// bool, string, numeric
	var a bool = true
	var b string = "Benji"
	var c rune = 'a' // es un int32 y unicode 97
	// uint dependera del so de 32 o 64 bits -> uint32/uint64
	var e uint8 = 255
	var d float64 = 23.56

	// Mantener logica de variables
	_ = 123
	var _ string = "valor de prueba"

	// Valores cero
	var x string
	var y uint
	var z bool

	// fmt.Printf para formateo
	// El verbo %T = tipo, %v=cualquier valor
	fmt.Printf("Tipo: %T, Valor: %v\n", a, a)
	fmt.Printf("Tipo: %T, Valor: %v\n", b, b)
	fmt.Printf("Tipo: %T, Valor: %v\n", c, c)
	fmt.Printf("Tipo: %T, Valor: %v\n", d, d)
	fmt.Printf("Tipo: %T, Valor: %v\n", e, e)

	fmt.Printf("%q\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)

	var age = 70
	fmt.Println("adulto:", age >= 18 && age <= 60)
}
