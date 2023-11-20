package storage

import (
	"mini_proyecto/shared"
)

func (s Storage) GetEmployee(id string) ( /*domain.Employee*/ []shared.Employee, error) {
	// Unico método que implementa la interfaz EmployeeStorage que nos va a devolver un empleado.

	rows, err := s.Col.Query("SELECT * FROM employees WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	// Iterate Rows
	var users []shared.Employee
	for rows.Next() {
		var user shared.Employee
		err := rows.Scan(&user.Id, &user.Name, &user.Status)
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, shared.ErrNotFound
		// Aquí se regresa este error centinela para que el caller sepa que no se encontró el documento.
		// Este error existe en nuestro dominio. El objetivo es comunicar diferentes tipos de error entre capas de la aplicación sin tener que importar paquetes de otras capas en este caso Mongo. Si devolvieramos el error tal cual como lo retorna Mongo, estaríamos exponiendo la implementación de la capa de almacenamiento a la capa de dominio. Y si en algún momento decidimos cambiar la capa de almacenamiento, tendríamos que cambiar la capa de dominio también. Es un error contra el principio de clean architecture.
	}

	return users, nil
}
