// Utilizacion de channel con buffer para limitar la cantidad de procesos concurrenten que se estan ejecutando

package main

import (
	"fmt"
	"sync"
	"time"
)

// En el primer for los espacios del channel estan vacios. Entonces la goRutine1 ocupa la primera posición 0 del channel.
// Luego la goRutine2 ocupa la posición 1 del channel.
//Cuando se va a ejecutar la 3ra goRutine ve que todos los espacios estan ocupados y espera.
//c := [][]
//c := [goRutine1][]
//c := [goRutine1][goRutine2]
//c := [][goRutine2]
//c := [goRutine3][goRutine2]

func main() {
	c := make(chan int, 2) // Lo mejor es delimitar la cantidad de espacios para no colapsar la maquina de procedimientos ejecutandose en paralelo.
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1 // Aquí lo que se esta haciendo es darle un cupo al canal que se crea en la linea 12 a la subrutina que se esta ejecutando en la linea 17.
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) { // Emulando una petición a una base de datos.
	defer wg.Done()
	fmt.Println("ID", i)
	time.Sleep(4 * time.Second)
	fmt.Println("ID ", i, "finished")
	<-c // Aquí lo que hace es liberar el cupo en el canal
}
