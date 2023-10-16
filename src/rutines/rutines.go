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

func Example3() {
	for i := 0; i < 5; i++ {
		select { // Es la sentencia que nos permite leer de varios canales. Donde por cada "case" tendremos un canal diferente.
		case <-time.After(1 * time.Second): // El canal time.After() nos devuelve un canal que se ejecuta después de un tiempo determinado.
			fmt.Println("Doing something...")
		}
	}
	fmt.Println("Operation completed")

	// Main
	// go SimulateLongOperation() // Aquí la go rutine
	// time.Sleep(6 * time.Second)
}
