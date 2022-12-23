package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {

	e := Employee{}

	e.id = 69
	e.name = "Vic Flores"

	fmt.Println("Employee", e)
}
