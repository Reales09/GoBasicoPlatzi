package main

import "fmt"

type pc struct {
	ram   int
	disco int
	brand string
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de disco y es una %s", myPc.ram, myPc.disco, myPc.brand)
}

func main() {

	myPc := pc{ram: 16, disco: 1000, brand: "Lenovo"}

	fmt.Println(myPc)

}
