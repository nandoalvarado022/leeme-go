// https://www.youtube.com/watch?v=3BK1c__EUig
// En este mini proyecto manejaremos una arquitectura limpia, hexagonal o DDD.
// https://github.com/jairogloz/senior-go-projects/tree/master/testing/003_mocks_self_made

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"mini_proyecto/service"
	"mini_proyecto/storage"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	const dbHost = "localhost"
	const dbPort = "3306"
	const dbPassword = "12345678"
	const dbUser = "root"
	const dbName = "test"

	DSN := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName // Data source name
	// mongoURI := os.Getenv("MONGO_URI")

	mysqlClient, disconnectFunc, err := storage.NewMysqlClient(DSN)
	if err != nil {
		panic(err)
	}

	defer disconnectFunc()

	employeeStorage := storage.NewEmployeeStorage(mysqlClient) // No se en que caso se utiliza cuando utilizamos mysql. En mongo si se utiliza para guardar la referencia a la colección.

	employeeService := service.NewService(employeeStorage)
	fmt.Println("employeService: ", employeeService)

	var employees, _ = employeeService.ValidateEmployee("1234")

	fmt.Println("employees: ", employees)
	/*
		ids := []string{
			// "6538aa4995230dd3e49726f1",  // valid
			"655a7771402c7d1db373a8b8", // invalid
			// "64d7b8698029c0ab12a2cd1a",  // not found
			// "64d7b8698029c0ab12a2cd1g_", // error converting id to object id
		}

		for _, id := range ids {
			isValid, err := employeeService.ValidateEmployee(id)
			fmt.Printf("id: %s, isValid: %v, err: %v\n", id, isValid, err)
		}
	*/
}
