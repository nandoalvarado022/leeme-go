// Go lenguaje concurrente compilado

// VIDEOS:
// https://www.youtube.com/watch?v=PNBVh9ILLiY (Patron de diseño factory y conexion a BD mysql y Postgres)

// TERMINOS
// Funciones variadicas -> Son funciones que reciben un numero variable de parametros. Se reciben así func sum (params ...int) aquí sum se puede llamar con multiples
// parametros.
//  Receiver functions -> es la forma que tiene GO para implementar métodos a los structs
// Zero values: valores por defecto que se le asignan a las variables caundo se declaran pero no tienen asignacion. Ejemplo: int = 0, string = "", bool = false
// COMANDOS
// go build src/main.go -> compila el archivo main.go.Necesario para tomar los cambios que  hagamos en el archivo main.go
// go run src/main.go -> compila y ejecuta el archivo main.go

// MODULOS
// go mod init example.com/leeme_go // Crea un archivo go.mod que es el archivo de dependencias del proyecto. Es una caracteristica en GO que permite que los proyectos sean independientes de la ubicacion en la que se encuentren.
// Tambien podria ser go mod init github.com//nestor94/testingmodule
// En el siguiente ejemplo utilizamos un alias para la utilizacion de un mismo modulo pero otra version
// import (
// 	"github.com/donvito/helloworld"
// 	hello2 "github.com/donvito/helloworld/v2"
// )

// func main(){
// 	helloworld.SayHello()
// 	hello2.SayHello()
// }

// Deferenciar es viajar a la direccion de memoria a la que apunta
//  Rutinas en GO: Son funciones que se ejecutan en paralelo. Se crean con la palabra reservada go.
// Cuando tienes un PC con 2 nucleos esto quiere decir que GO puede correr dos rutinas simultaneamente.
// La palabra go crea concurrencia instantanea sin tener que crear hilos o procesos. Esto es una de las ventajas de GO
// Concurrencia: Es la composición de varias tareas independientes que se ejecutan al mismo tiempo. Ejemplo: 2 personas hablando por telefono en el mismo telefono
// Paralelismo: es la ejecución simultánea de varias tareas (o varias partes de una única tarea) en distintos procesadores. Ejemplo: 2 personas hablando por telefono al mismo tiempo pero en dos telefonos diferentes
// var wg = sync.WaitGroup acumula un conjunto de runtimes y las va liberando poco a poco
// Deadlocks: En un deadlock hay una goroutine que está esperando leer o escribir en un canal, sin embargo ya no hay ninguna goroutine ejecutándose pues están esperándose las unas a las otras; están en un punto muerto del que no se puede avanzar.
/* Los canales en GO:
// - Son sincronos
// https://www.youtube.com/watch?v=Rh3eSyd67h0
Channels: medios de comunicacion entre diferentes rutinas
Buffer channels: Son canales que tienen un buffer. El buffer es un espacio en memoria que se le asigna al canal para que pueda almacenar datos. Si el canal no tiene buffer, el canal no puede almacenar datos y se bloquea hasta que se lea un dato del canal.
*/
// Instalar librerias de terceros: go get golang.org/x/tour
// Variables de enforno: se utiliza la libreria
// go mod edit -replace
// go mod edit -dropreplace github.com/labstack/echo // Deshace el replace anteriormente hecho
// go.sum // Es un archivo que contiene las versiones de las dependencias que se estan usando en el proyecto
// go mod tidy // Actualiza las dependencias del proyecto. Elimina de go.mod las dependencias que no se estan usando en el proyecto
// GO no tiene manejo de excepciones por lo cual siempre hay que jugar con la variables err
// Para exportar una variable o función se debe poner la primera letra en mayuscula

package main

import (
	"fmt"
	pk "leeme_go/src/general" // Se importa el paquete mypackage y se le asigna el alias pk
	"leeme_go/src/rutines"
	"leeme_go/src/stringers"
	"strconv" // Se importa el paquete strconv que sirve para convertir un string a un int
	"time"
)

type car struct { // Se crea una estructura. Es como una clase en otros lenguajes
	brand string
	year  int
}

func publicar(canal chan<- int) { // El canal solo va a enviar datos
	canal <- 7 // Se envia el valor 7 por el canal
	fmt.Println("Final de la func publicar")
}

func escuchar(canal <-chan int) { // El canal solo va a recibir datos
	i := <-canal   // Llega aquí y espera a que el canal se llene
	fmt.Println(i) // Luego de que pasa por publicar se imprime el valor que se envio por el canal
}

func canales() {
	canal := make(chan int, 20) // Se crea un canal.
	go escuchar(canal)
	publicar(canal)
}

func primeraFuncion(num1, num2 int, c string) string { // Aquí se especifica que num1 y num2 son de tipo int y c de tipo string y que la función retorna un string
	suma := num1 + num2
	mensaje := fmt.Sprintf("%s es el mensaje y la suma es: %d", c, suma)
	return mensaje
}

func doubleReturn(a int) (c, d int) { // Aqui se especifica que la función retorna dos valores de tipo int
	return a, a * 2
}

func keywords() {
	// Defer -> se ejecuta al final de la función
	// Se usa para cerrar archivos, cerrar conexiones a base de datos, cerrar conexiones a servidores, etc.
	defer fmt.Println("Hola") // Se ejecuta al final de la función
	fmt.Println("Mundo")

	// Contiue y break
	// Se usa para salir de un ciclo o para saltar una iteración
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue // Salta la iteración
		}
		fmt.Println(i)
		if i == 8 {
			break // Sale del ciclo
		}
	}
}

func arrayAndSlice() {
	// Array
	var array [4]int // Al se un valor inmutable como son los arrays en GO no podemos agregar otro elemento
	// Si podemos cambiar el valor de un elemento
	array[0] = 1
	fmt.Println(array)
	fmt.Println(len(array)) // len() -> retorna el tamaño del array
	fmt.Println(cap(array)) // cap() -> retorna la capacidad del array

	// Slice
	slice := []int{1, 2, 3, 4} // Al ser un valor mutable podemos agregar más elementos
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en los slices
	fmt.Println(slice[0])   // Acceder a un elemento del slice
	fmt.Println(slice[:3])  // Acceder a los elementos del slice desde el inicio hasta el indice 3
	fmt.Println(slice[2:4]) // Acceder a los elementos del slice desde el indice 2 hasta el indice 4
	fmt.Println(slice[2:])  // Acceder a los elementos del slice desde el indice 2 hasta el final

	// Append
	// Agrega un elemento al slice
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...) // Aquí se agrega el slice newSlice al slice slice
	fmt.Println("New slice:", slice)
}

func privateAndPublic() {
	var myCar pk.CarPublic // Se crea una instancia de la estructura CarPublic que se encuentra en el paquete mypackage
	fmt.Println(myCar)
	pk.PrintMessage()
}

func structs() {
	myCar := car{brand: "Ford", year: 2020} // Se crea una instancia de la estructura car
	fmt.Println(myCar)

	// Otra forma de crear una instancia de la estructura car
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 2020
	fmt.Println(otherCar)
}

func diccionarios() {
	m := make(map[string]int) // Se crea un diccionario vacio. Aquí make es una función que crea un diccionario vacio
	m["Jose"] = 14
	m["Pedro"] = 20
	fmt.Println(m)

	// Iterar un diccionario
	for i, v := range m {
		fmt.Println(i, v) // i -> key, v -> value
	}

	// Encontrar un valor en el diccionario
	value := m["Jose"]
	fmt.Println(value)

	value, ok := m["Jose"]
	fmt.Println(value, ok) // ok retorna true si el valor existe y false si no existe
}

func isPalindrome(text string) {
	var textReverse string
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if textReverse == text {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func interactArrayAndSlice() {
	slice := []string{"hola", "que", "tal"}
	for i := range slice {
		fmt.Println(slice[i])
	}

	isPalindrome(("ana"))
}

func conditionalIteration() {
	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	value, err := strconv.Atoi("10") // Atoi convierte un string a un int
	if err != nil {                  // El nil es como el null de otros lenguajes.
		// nil se usa para saber si una variable tiene un valor o no y tambien para saber si una funcion retorno un valor o no
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	// Switch
	switch modulo := 4 % 2; modulo { // Aquí se declara una variable y se le asigna el valor de 4 % 2
	case 0:
		fmt.Println("Es par")
	}

	switch { // De esta manera se puede hacer un switch sin declarar una variable
	case 2 == 2:
		fmt.Println("Es par")
	}
}

func basic() {
	// Funciones con doble retorno
	fmt.Println("Funciones con doble retorno:")
	value1, value2 := doubleReturn(2) // Aquí se asigna el valor de retorno a las variables value1 y value2
	fmt.Println(value1, value2)

	// Tipos de datos
	fmt.Println("Tipos de datos:")
	const pi float64 = 3.14
	const saludo string = "Hola Mundo"
	const numero int = 123
	const booleano bool = true
	base := 12 // Aquí se asigna el tipo de dato automáticamente
	fmt.Println("Hello World", saludo, numero, booleano)
	fmt.Println(pi, base)

	// Saber el tipo de dato
	fmt.Printf("El tipo de dato de la variable numero es: %T\n", numero) // %T -> tipo de dato

	fmt.Println("____________Casting:_____________")
	myValue, err := strconv.ParseInt("10", 0, 64) // ParseInt convierte un string a un int
	fmt.Println(myValue, err)

	x := 10
	y := 50
	fmt.Println(x + y)

	// Obteneindo modulo
	fmt.Println("Modulo es:", x%y)

	// printF
	platzi := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", platzi, cursos) // %s -> string, %d -> int, %f -> float, %v -> cualquier tipo de dato

	// Sprintf -> retorna un string
	message := fmt.Sprintf("%s tiene más de %d cursos", platzi, cursos)
	fmt.Println(message)
}

func usingTime() {
	layout := "2006-01-02"
	date := "2022-02-01"
	time, _ := time.Parse(layout, date)
	formattedTime := time.Format("02 Jan 2006")
	fmt.Println(formattedTime)
}

func main() {
	// fmt.Println(returnsWithNames.GetValues(2))
	// interfaces.Basic()

	// pointers.PunterosyStructs()
	// pointers.UsingPlayerExample()

	// usingTime()

	rutines.Basic()
	rutines.Example2()

	// basic()

	// canales()

	stringers.Basic()

	// privateAndPublic()

	// structs()

	// diccionarios()

	// interactArrayAndSlice()

	// arrayAndSlice()

	// keywords
	// keywords()

	// Iteraciones condicionales
	// conditionalIteration()

	// Iteraciones
	// fmt.Println("Iteraciones:")
	// conditionalIteration()

	// Funciones
	// result := primeraFuncion(2, 4, "Hi there")
	// fmt.Println(result)

	// var myVar int
	// fmt.Println(myVar)
}
