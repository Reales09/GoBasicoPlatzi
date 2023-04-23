package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string

	text = strings.ToLower(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])

	}

	if textReverse == text {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {

	isPalindromo("casas")

	slice := []string{"hola", "que", "hace"}

	//Imprime solo el valor
	for _, valor := range slice {
		println(valor)
	}

	// Imprime el indice y el valor

	for i, valor := range slice {
		println(i, valor)
	}
	// imprime solo indice

	for i := range slice {
		println(i)
	}

	//ama
	// amor a roma

}
