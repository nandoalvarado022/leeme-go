package employee

import (
	"errors"
	"fmt"
	"mini_proyecto/shared"
)

func (s Service) ValidateEmployee(id string) (bool, error) {
	employee, err := s.storage.GetEmployee(id)

	if err != nil {
		if errors.Is(err, shared.ErrNotFound) {
			return false, fmt.Errorf("Employee with id '%s' not found", id) // Si aquí estuvieramos en la capa HTML deberiamos devolver un 404 de que no se encontró el recurso
		}
		return false, fmt.Errorf("Error getting employee: %w", err) // Aquí un 500
	}

	if employee[0].Status != "active" {
		return false, fmt.Errorf("Employee with id '%s' is not active", id)
	}

	fmt.Println(employee)

	return true, nil
}
