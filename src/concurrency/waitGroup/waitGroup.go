// waitGroup es una estructura que nos permite esperar a que un conjunto de rutinas termine su ejecución.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg) // Se envia como referencia para evitar que se cree una copia de la variable.
	}

	wg.Wait() // Espera a que el contador llegue a cero.
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Siempre se ejecuta al final de la función.
	// wg.Done lo que hace es decrementar el contador por uno.
	fmt.Print("Doing something ", i, "\n...")
	time.Sleep(2 * time.Second)
	fmt.Println("Operation completed")
}
