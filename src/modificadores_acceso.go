package main

import (
	"fmt"

	pk "src/mypackage"
)

func main() {

	var myCar = pk.CarPublic{}
	myCar.Brand = "Ferrari"
	myCar.Year = 2023

	// fmt.Println(pk.CarPublic{Brand: "Ford", Year: 2020})
	fmt.Println(myCar)

	pk.PrintMessage("Reales")
}
