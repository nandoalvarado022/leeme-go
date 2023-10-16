// En este ejemplo se cancela la go rutina después de 3 segundos utilizando el paquete context.

package main

import (
	"context"
	"fmt"
	"time"
)

func SimulateLongOperation(ctx context.Context) {
	for i := 0; i < 5; i++ {
		select { // Es la sentencia que nos permite leer de varios canales. Donde por cada "case" tendremos un canal diferente.
		case <-time.After(1 * time.Second): // El canal time.After() nos devuelve un canal que se ejecuta después de un tiempo determinado.
			fmt.Println("Doing something...")
		case <-ctx.Done(): // Aquí se cancela la ejecución de la go rutine
			fmt.Println("Operation canceled")
			return
		}
	}
	fmt.Println("Operation completed")
}

func main() {
	ctxHijo, _ := context.WithTimeout(context.Background(), 3*time.Second) // Aquí 3 seria el tiempo de espera. Si se pasa de ese tiempo, se cancela la ejecución de la go rutine
	go SimulateLongOperation(ctxHijo)
	time.Sleep(6 * time.Second) // Aquí espera 6 segundos para que se ejecute la go rutine
}
