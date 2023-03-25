package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time is: ", now)

	formatDate := time.Date(2022, time.Now().Month(), 5, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go was launched at: ", formatDate)
	fmt.Println(formatDate.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)
	fmt.Printf("Tipo: %T, Valor: %v\n", now.Format("2006-01-02 15:04:05"), now.Format("2006-01-02 15:04:05"))
}
