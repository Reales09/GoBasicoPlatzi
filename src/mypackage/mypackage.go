package mypackage

import "fmt"

// CarPublic Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage Imprime un mensaje
func PrintMessage(nombre string) {
	fmt.Println("Hola: ", nombre)
}

func printMessage1(nombre string) {
	fmt.Println("Hola: ", nombre)
}
