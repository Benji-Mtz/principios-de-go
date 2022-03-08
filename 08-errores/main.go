package main

import (
	"errors"
	"fmt"
)

func main() {
	// content, err := ioutil.ReadFile("C:/Users/Benji/Documents/Go/src/github.com/Benji-Mtz/ED-Bases/08-errores/hello.txt")
	// if err != nil {
	// 	fmt.Printf("Ocurrio un error: %v", err)
	// 	return
	// }
	// fmt.Println(string(content))

	result, err := Division(10, 0)

	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	}

	fmt.Println(result)
}

func Division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir entre cero")
	}

	return dividendo / divisor, nil
}
