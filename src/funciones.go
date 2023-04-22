package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {

	return a, a * 3
}

func main() {
	normalFunction("Some message")
	tripleArgument(1, 2, "Some message")
	value := returnValue(2)
	println("value", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("value1", value1)

	fmt.Println("value2", value2)

	//value1, _ := doubleReturn(2) // Si no se necesita un valor de retorno se puede omitir
}
