package main

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text

}

func main() {
	c := make(chan string, 2)
	c <- "Hola"
	c <- "Mundo"

	fmt.Println(len(c), cap(c))

	// Range y close

	close(c)

	// c <- "Mensaje 3"

	for message := range c {
		fmt.Println(message)
	}

	// select

	email1 := make(chan string)
	email2 := make(chan string)

	go message("Mensaje 1", email1)
	go message("Mensaje 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-email1:
			fmt.Println("Email recibido de email1: ", msg1)
		case msg2 := <-email2:
			fmt.Println("Email recibido de email2: ", msg2)
		}
	}

}
