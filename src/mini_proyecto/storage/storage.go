package storage

import (
	"database/sql"
	"mini_proyecto/domain"
	"mini_proyecto/shared"
)

// Storage implements the domain.EmployeeStorage interface
type Storage struct {
	// Col *mongo.Collection // Aquí se guarda la referencia a la colección de mongo para futuras referencias.
	Get func(id string) ([]shared.Employee, error)
	Col *sql.DB
}

func NewEmployeeStorage(mysqlClient *sql.DB) domain.EmployeeStorage {
	// No se en que caso se utiliza cuando utilizamos mysql. En mongo si se utiliza para guardar la referencia a la colección.
	return &Storage{
		// Col: mysqlClient.Database("hr").Collection("employees"),
		Col: mysqlClient,
		Get: GetEmployee,
	}
}

func GetEmployee(id string) ([]shared.Employee, error) {
	return []shared.Employee{
		{
			Id:     1234,
			Name:   "John",
			Status: "Doe",
		},
	}, nil
}
