package main

import "fmt"

type rectangulo struct {
	base   float64
	altura float64
}

type cuadrado struct {
	base float64
}

type figuras2D interface { // Se crea una interfaz. Se diferencia de los struct porqu elas interface pueden tener funciones asociadas (receiver functions). Ademas las interfaces no pueden tener propiedades
	area() float64 // Se crea un método que va a retornar un float64
}

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
	endDate string
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "FT Employee, Hi there " + ftEmployee.name
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temp Employee, Hi there " + tEmployee.name
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func (r rectangulo) area() float64 {
	// Va a exportar un float64. Va a recibir un r de tipo rectangulo y a la función le llamaremos area.
	// area() acá realmente no seria una funcion sino un método, por esto es que se le pasa como parametro el tipo de dato al que va a pertenecer el método.
	return r.base * r.altura
}

func (c cuadrado) area() float64 { // Va a exportar un float64. Va a recivit un c de tipo cuadrado y le llamaremos area.
	return c.base * c.base
}

func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func main() {
	// Metodo en el cual se puede compartir diferenets otros métodos
	// myCuadrado := cuadrado{base: 2}
	// myRectangulo := rectangulo{base: 2, altura: 4}
	// fmt.Println(myCuadrado.area())
	// fmt.Println(myRectangulo.area())
	// calcular(myCuadrado)
	// calcular(myRectangulo)

	// Lista de interfaces aceptando diferentes tipos de datos
	// myInterface := []interface{}{"Hola", 12, 4.90}
	//fmt.Println(myInterface...)

	Employees()
}

func Employees() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Pedro"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
	//GetMessage(ftEmployee)
}
