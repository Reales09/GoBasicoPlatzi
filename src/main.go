package main

import "fmt"

func main() {

	//Declaración de constantes

	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaración de variables enteras
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es: ", areaCuadrado)

	x := 10
	y := 50

	//Suma

	result := x + y

	fmt.Println("El resultado de la suma es: ", result)

	//Resta

	result = y - x

	fmt.Println("El resultado de la resta es: ", result)

	//Multiplicación

	result = x * y

	fmt.Println("El resultado de la multiplicación es: ", result)

	//División

	result = y / x

	fmt.Println("El resultado de la división es: ", result)

	//Módulo

	result = y % x

	fmt.Println("El resultado del módulo es: ", result)

	//Incremento

	x++

	fmt.Println("El valor de x es: ", x)

	//Decremento

	y--

	fmt.Println("El valor de y es: ", y)

	//Retos
	// -Area de un triángulo, trapecio y círculo

	//Area del triángulo
	base1 := 10   //base del triángulo
	altura1 := 15 //altura del triángulo

	areaTriangulo := (base1 * altura1) / 2

	fmt.Println("El área del triángulo es: ", areaTriangulo)

	//Area del trapecio

	baseMayor := 10
	baseMenor := 5
	alturaTrapecio := 15

	resultadoTrapecio := ((baseMayor + baseMenor) * alturaTrapecio) / 2

	fmt.Println("El área del trapecio es: ", resultadoTrapecio)

	//Area del círculo

	const pi3 = 3.1416

	var radio float64 = 10

	resultadoCirculo := pi3 * (radio * radio)

	fmt.Println("El área del círculo es: ", resultadoCirculo)

}
