package main

import "fmt"

type A struct{}

func (a A) Hi() string {
	return "A says Hi "
}

type B struct{}

func (b B) Hi() string {
	return "B says Hi"
}

type Greeter struct {
	A
	B
}

func (g Greeter) Speak() string {
	// return g.Hi() —> This method belongs to A or B?
	return g.A.Hi() + g.B.Hi()
}

func main() {
	saludo := Greeter{}
	fmt.Println(saludo.Speak())
}
