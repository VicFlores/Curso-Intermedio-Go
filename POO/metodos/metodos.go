package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) setId(id int) {
	e.id = id
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) getId() int {
	return e.id
}

func (e *Employee) getName() string {
	return e.name
}

func main() {

	e := Employee{}

	e.id = 69
	e.name = "Vic Flores"
	fmt.Println("Employee", e)

	e.setId(777)
	fmt.Println("New ID employee", e.id)

	e.setName("Ferman")
	fmt.Println("New name employee", e.name)

	fmt.Println("Get id", e.getId())
	fmt.Println("Get name", e.getName())
}
