package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())

}

func main() {
	myCuadrado := cuadrado{base: 10}

	myRectangulo := rectangulo{base: 15, altura: 10}

	calcular(myCuadrado)

	calcular(myRectangulo)

	// Lista de interfaces
	myInterfaces := []interface{}{"Hola", 12, 3.1416, true}

	// Impresion de cada elemento de la lista
	fmt.Println(myInterfaces...)

	// Impresion de la lista
	fmt.Println(myInterfaces)

	//Impresion del primer elemento de la lista
	fmt.Println(myInterfaces[0])

}
