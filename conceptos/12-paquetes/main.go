package main

import (
	"fmt"

	"github.com/Benji-Mtz/principios-de-go/conceptos/12-paquetes/greet"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello")
	fmt.Println(greet.English())
	fmt.Println(greet.Italian())
	fmt.Println(greet.Spanish())
	fmt.Println(quote.Hello())
}
