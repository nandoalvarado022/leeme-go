/*
Los channels son un tipo de dato en GO que nos permite comunicar diferentes rutinas entre si. Es como un tubo por donde se envian mensajes
Los channels pueden ser de 2 tipos: Unbuffered y Buffered
Unbuffered channel: Espera una función o una rutina para recibir el mensaje, es bloqueada por ser llamada en la misma función
Buffered channel: Se puede llamar de manera inmediata, en el siguiente ejemplo 2 es el numero de canales que pueden ser usados
Se les puede establecer un tamaño a los canales. Por ejemplo: c := make(chan int, 2) // Aquí 2 es el numero de canales que pueden ser usados

Los canales en GO:
Son sincronos
https://www.youtube.com/watch?v=Rh3eSyd67h0
Channels: medios de comunicacion entre diferentes rutinas
Buffer channels: Son canales que tienen un buffer. El buffer es un espacio en memoria que se le asigna al canal para que pueda almacenar datos. Si el canal no tiene buffer, el canal no puede almacenar datos y se bloquea hasta que se lea un dato del canal.

Canal de solo escritura: Flecha <- a la derecha de chan
Ejemplo: func Generator(c chan<- int)

Canal de solo lectura: Flecha <- a la izquierda de chan
Ejemplo: func Print(c <-chan int)

*/

package main

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
	type car struct {
		brand string
		year  int
	}
	myCar := car{
		brand: "Ford",
	}
	fmt.Println("Starting Example2...")
	c := make(chan int)
	go doSomething(c)
	<-c
	if myCar.year != nil {
		fmt.Println("Existe year")
	} else {
		fmt.Println("No existe year")
	}
	fmt.Println("Terminó")
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

func main() {
	// Basic()
	Example2()
	// Example3()
}
