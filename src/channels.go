package main

import (
	"fmt"
)

func say(text string, c chan<- string) {
	c <- text
}

func main() {

	// c := make(chan string) // Channerl de forma dinámica

	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("world", c)

	fmt.Println(<-c)

}
