package storage

import (
	"database/sql"
	"mini_proyecto/pkg/domain/employee"
	"mini_proyecto/shared"
)

// Storage implements the domain.EmployeeStorage interface
type Storage struct {
	// Col *mongo.Collection // Aquí se guarda la referencia a la colección de mongo para futuras referencias.
	Otro        func(id string) ([]shared.Employee, error)
	mysqlClient *sql.DB
}

func NewEmployeeStorage(mysqlClient *sql.DB) employee.EmployeeStorage {
	// Para iniciarlizarlo solo necesitamos un cliente que este conectado a al BD en este caso mysqlClient.
	// No se en que caso se utiliza cuando utilizamos mysql. En mongo si se utiliza para guardar la referencia a la colección.
	return &Storage{
		mysqlClient: mysqlClient, // En el caso de mongo aquí lo se hace es:mongoClient.Database("hr").collection("Employees")
	}
}
