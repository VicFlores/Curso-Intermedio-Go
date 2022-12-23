package main

import "testing"

func TestGetFullTimeEmployeeByID(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "abc12345",
			mockFunc: func() {
				GetEmployeeByID = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "ABC12345",
						Name: "Vic Flores",
						Age:  21,
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					DNI:  "ABC12345",
					Name: "Vic Flores",
					Age:  21,
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeByID
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()

		ftEmployeeById, err := GetFullTimeEmployeeByID(test.id, test.dni)

		if err != nil {
			t.Error("Error when getting employee")
		}

		if ftEmployeeById.Age != test.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ftEmployeeById.Age, test.expectedEmployee.Age)
		}

	}

	GetEmployeeByID = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
