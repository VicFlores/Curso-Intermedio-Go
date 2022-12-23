package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(dni string) (Person, error) {
	/* Simulando un llamado a una DB */
	time.Sleep(5 * time.Second)

	return Person{}, nil
}

var GetEmployeeByID = func(id int) (Employee, error) {
	/* Simulando un llamado a una DB */
	time.Sleep(5 * time.Second)

	return Employee{}, nil
}

func GetFullTimeEmployeeByID(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	/* Get employee */

	getEmployee, err := GetEmployeeByID(id)

	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Employee = getEmployee

	/* Get person */

	getPerson, err := GetPersonByDNI(dni)

	if err != nil {
		return ftEmployee, err
	}

	ftEmployee.Person = getPerson

	return ftEmployee, nil

}
