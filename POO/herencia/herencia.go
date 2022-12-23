package main

import "fmt"

type Person struct {
	age  int
	name string
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.id = 132
	ftEmployee.name = "Vic"
	ftEmployee.age = 22
	fmt.Println("Full Time Employee", ftEmployee)

}
