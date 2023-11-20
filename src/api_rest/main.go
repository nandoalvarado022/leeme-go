// https://www.youtube.com/watch?v=pQAV8A9KLwk
// Run: go run src/api_rest/main.go

package api_rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all tasks")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create task")
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}
	newTask.ID = len(tasks) + 1
	json.Unmarshal(reqBody, &newTask) // Asignando el reqBody a la variable newTask. Estamos pasando de un arreglo de bytes a un objeto de tipo task.
	tasks = append(tasks, newTask)
	w.WriteHeader(http.StatusCreated)                  // Para decirle que se ha creado un nuevo recurso
	w.Header().Set("Content-Type", "application/json") // Para decirle que el contenido que va a devolver es de tipo json. Sirve para que por ejemplo en postman sepa que el contenido que va a devolver es de tipo json y lo muestre pretty.
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // Eliminando el elemento del slice. Con el :i se añade los elementos que esten antes del indice i y con el i+1 se añaden los elementos que esten despues del indice i.
			fmt.Fprintf(w, "The task with ID %v has been removed successfully", id)
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask task
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
		return
	}

	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // Eliminando el elemento del slice. Con el :i se añade los elementos que esten antes del indice i y con el i+1 se añaden los elementos que esten despues del indice i.
			updatedTask.ID = id
			tasks = append(tasks, updatedTask)
			fmt.Fprintf(w, "The task with ID %v has been updated successfully", id)
		}
	}
}

func Main() {
	fmt.Println("Hello, playground 3")
	router := mux.NewRouter().StrictSlash(true)

	// Task routes
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
