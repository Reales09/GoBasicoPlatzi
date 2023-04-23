package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pedro"] = 20
	m["Juan"] = 30

	fmt.Println(m)

	// Recorrer map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	// si no exiten devuelve 0 y false
	value, ok := m["Josep"]
	fmt.Println(value, ok)

	// si exiten devuelve el valor y true
	value1, ok := m["Jose"]
	fmt.Println(value1, ok)

}
