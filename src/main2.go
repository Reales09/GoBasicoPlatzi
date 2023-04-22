package main

import "fmt"

func main() {
	//declaración de variables

	helloMesaage := "Hello"
	worldMessage := "world"

	//Println = función para imprimir en consola
	fmt.Println(helloMesaage, worldMessage)
	fmt.Println(helloMesaage, worldMessage)

	//printf = función para imprimir con formato

	nombre := "Platzi"
	cursos := 500

	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	//Sprintf = devuelve el valor formateado

	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)

	fmt.Println(message)

	//Tipo de datos

	fmt.Printf("helloMessage: %T\n", helloMesaage)
	fmt.Printf("cursos: %T\n", cursos)

}
