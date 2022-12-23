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
	endDate string
}

func (f FullTimeEmployee) getMessage() string {
	return "Full time employee"
}

type TemporalEmployee struct {
	Person
	Employee
	taxtRate int
}

func (t TemporalEmployee) getMessage() string {
	return "Temporal employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println("Get message", p)
}

func main() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.id = 132
	ftEmployee.name = "Vic"
	ftEmployee.age = 22
	fmt.Println("Full Time Employee", ftEmployee)

	tEmployee := TemporalEmployee{}

	getMessage(ftEmployee)
	getMessage(tEmployee)

}
