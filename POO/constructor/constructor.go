package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {

	/* Forma 1 de instanciar */
	e := Employee{}
	fmt.Println("Employee #1", e)

	/* Forma 2 de instanciar */
	e2 := Employee{id: 1, name: "Vic", vacation: true}
	fmt.Println("Employee #2", e2)

	/* Forma 3 de instanciar */
	e3 := new(Employee)
	fmt.Println("Employee #3", *e3)

	/* Forma 4 de instanciar */
	e4 := NewEmployee(2, "Ferman", true)
	fmt.Println("Employee #4", *e4)

}
