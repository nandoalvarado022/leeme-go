package main

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan int, results chan<- int) { //
	fmt.Println("Worker with id", id, "started")
	for j := range jobs {
		fmt.Printf("Worker with id %v started job %v\n", id, j)
		results <- Fibonacci(j)
		fmt.Printf("Worker with id %v finished job %v\n", id, j)
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	tasks := []int{5, 4, 3} // Estos son los valores de los que queremos calcular los fibonacci
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))
	for w := 1; w <= 3; w++ { // Utilizamos 3 workers
		go Worker(w, jobs, results)
	}
	time.Sleep(3 * time.Second)
	for _, task := range tasks {
		println("Sending task", task)
		jobs <- task
		time.Sleep(1 * time.Second)
	}
	close(jobs)
	for a := 1; a <= len(tasks); a++ {
		fmt.Println(<-results)
	}
}
