package main

import "fmt"

type pc struct {
	ram   int
	disco int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")

}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2

}

func main() {
	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disco: 1000, brand: "Lenovo"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)

	myPC.duplicateRAM()

	fmt.Println(myPC)

	myPC.duplicateRAM()

	fmt.Println(myPC)

}