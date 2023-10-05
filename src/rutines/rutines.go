package rutines

import (
	"fmt"
	"time"
)

func Basic() {
	fmt.Println("Hi there!")
	time.Sleep(5 * time.Second)
	fmt.Println("Bye there!")
}

func Example2() {
	c := make(chan int)
	go doSomething(c)
	<-c
}

func doSomething(c chan int) { // Func utilizada para crear una rutina
	time.Sleep(3 * time.Second) // Debe esperar 3 segundos para ejecutar la siguiente linea de código
	fmt.Println("Done")
	c <- 1
}
