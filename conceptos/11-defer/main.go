package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error al crear el archivo: %v", err)
		return
	}

	// siempre ejecutarra file.close al final del programa y barre donde se use
	defer file.Close()

	_, err = file.Write([]byte("Hello Benji"))
	if err != nil {
		// file.Close()
		fmt.Printf("Ocurrió un error al crear el archivo: %v", err)
		return
	}

	file.Close()
}
