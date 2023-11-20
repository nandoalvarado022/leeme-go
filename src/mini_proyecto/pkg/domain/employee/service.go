package employee

import (
	"fmt"
)

type Service struct { // Estructura que es la que va a implementar la interface EmployeeService.
	// mysqlClient *sql.DB
	storage          EmployeeStorage
	mostrandoMensaje func() string
}

func NewService(storage EmployeeStorage) EmployeeService { // Aquí estamos pasando una implementación de la interface EmployeeStorage.
	// La cual asignamos a la variable storage para poder usarla luego en los métodos de la estructura Service (ahora solo tendriamos validateEmployee).
	// Esto es lo que se conoce como inyección de dependencias.
	if storage == nil {
		panic("storage is nil")
	}

	return &Service{
		storage: storage,
		mostrandoMensaje: func() string {
			fmt.Println("Obteniendo clientes")
			return "Obteniendo clientes"
		},
	}
}
