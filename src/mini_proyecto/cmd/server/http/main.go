// la carpeta cmd contiene los archivos de entrada para la aplicacion. En este caso solo tendremos el archivo main.go.
// main.go es quien se encarga de iniciar el servidor HTTP  utilizando la configuración proveniente de la carpeta pkg/api/config/.

// https://www.youtube.com/watch?v=3BK1c__EUig

// En este mini proyecto manejaremos una arquitectura limpia, hexagonal o DDD:
// https://github.com/jairogloz/senior-go-projects/tree/master/testing/003_mocks_self_made

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mini_proyecto/pkg/domain/employee"
	"mini_proyecto/pkg/storage"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func getEmployees(w http.ResponseWriter, r *http.Request, employeeService employee.EmployeeService) {
	vars := mux.Vars(r)
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	var employees, _ = employeeService.ValidateEmployee(id)
	fmt.Println("employees: ", employees)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func main() {
	fmt.Println("Hello, playground 1")
	router := mux.NewRouter().StrictSlash(true)

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

	mysqlClient, disconnectFunc, err := storage.NewMysqlClient(DSN)
	if err != nil {
		panic(err)
	}

	defer disconnectFunc()

	employeeStorage := storage.NewEmployeeStorage(mysqlClient) // No se en que caso se utiliza cuando utilizamos mysql. En mongo si se utiliza para guardar la referencia a la colección.

	employeeService := employee.NewService(employeeStorage)
	fmt.Println("employeService: ", employeeService)

	// API routes
	router.HandleFunc("/", indexRoute)
	// router.HandleFunc("/employees", getEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		getEmployees(w, r, employeeService)
	})
	log.Fatal(http.ListenAndServe(":3000", router))
}
