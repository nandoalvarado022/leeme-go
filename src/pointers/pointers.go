package pointers

import "fmt"

// Punteros: Es una dirección de memoria. Se usa para acceder a un valor en memoria. Se usa el operador
// Se usa * para obtener el valor de la variable a la que apunta el puntero
// Se usa & para obtener la dirección de memoria de una variable y el operador
// Recibir un parametro como puntero cuando:
//	1. Cuando se necesita que el método modifique al reciver y que esta modificación sea permanente
//	2. Cuando el receiver es muy grande y no queremos copiarlo. Para este ejemplo supongamos que la propiedad Records del struct Player tiene muchos elementos (15 MB)
// De esta manera declaramos un puntero:var p* int

type Pc struct { // Se crea una estructura. Es como una clase en otros lenguajes
	Ram   int
	Disk  int
	Brand string
}

type Player struct {
	Points  int
	Records []string
}

func (r *Player) AddPoints(points int) {
	// Al llamar este método nos damos cuenta que la direccion en memoria de la variable r es
	// diferente a la direccion de memoria de player1 impresa en el método UsingPlayerExample. Por esto al querer sumar los puntos
	// no se modifica el valor de player1. Esto se debe a que en GO se pasan los parametros por valor y no por referencia.
	// Para resolver esto se debe usar un puntero en el receiver function. Así: (r *Player)
	fmt.Printf("Player address: %p\n", &r) // Aquí se imprime la dirección de memoria de la variable r
}

func (myPC Pc) ping() { // Se crea un receiver function (método) para la estructura Pc.
	fmt.Println(myPC.Ram, " Pong")
	myPC.duplicateRAM()
	fmt.Println(myPC.Ram, " Pong")
}

func (myPC *Pc) duplicateRAM() {
	myPC.Ram *= 2
}

func PunterosYStructs() {
	// Los punteros en GO se utilizan para modificar propiedades de los strucs en GO de forma más eficiente

	a := 5
	b := &a // Aquí se le asigna a b la dirección de memoria de a
	// Una direccion de memoria es un numero hexadecimal
	fmt.Println(a, b)
	fmt.Println(*b) // Aquí se imprime el valor de la variable a la que apunta b
	*b = 100
	fmt.Println(a) // Aquí se imprime el valor de a que ahora es 100

	myPc := Pc{Ram: 16, Disk: 200, Brand: "MSI"}
	fmt.Println(myPc)

	myPc.ping()
}

func UsingPlayerExample() {
	player1 := Player{
		Points: 0,
	}

	player1.AddPoints(100)
	fmt.Printf("Player address: %p\n", &player1)
}
