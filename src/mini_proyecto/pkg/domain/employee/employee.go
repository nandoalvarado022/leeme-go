// En la carpeta dominio tenemos los tipos principales que son nativos de nuestro negocio.

package employee

import "mini_proyecto/shared"

type EmployeeService interface {
	// Es la que va a proveer los servicios de la aplicacion.
	//Actualmente solo tendremos ValidateEmployee que nos va a decir si el empleado es valido.
	ValidateEmployee(id string) (bool, error)
}

type EmployeeStorage interface {
	// Tambien tendremos una interface EmployeeStorage que tiene un unico método GetEmployee que nos va a devolver un empleado.
	GetEmployee(id string) ([]shared.Employee, error) // Esta función se utiliza dentro del validateEmployee
}
