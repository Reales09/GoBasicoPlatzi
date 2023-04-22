package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("El valor es 1")
	} else {
		fmt.Println("El valor no es 1")
	}

	//With and

	if valor1 == 1 && valor2 == 2 {
		fmt.Println("El valor es 1 y 2")
	} else {
		fmt.Println("El valor no es 1 y 2")
	}

	// With or

	if valor1 == 1 || valor2 == 2 { //With or
		fmt.Println("El valor es 1 o 2")
	} else {
		fmt.Println("El valor no es 1 o 2")
	}

	//Convertir texto a numero

	value, err := strconv.Atoi("hola")

	if err != nil {
		log.Fatal(err)
		fmt.Println(value)
	} else {
		fmt.Println("Value: ", value)
	}

}
