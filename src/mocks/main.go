package mocks

type Employee struct {
	id  int
	dni string
}

func GetEmployee(id int, dni string) (Employee, error) {
	return &Employee{
		id:  id,
		dni: "12345678",
	}
}
