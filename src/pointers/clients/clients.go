package clients

import "fmt"

type Cliente struct {
	nombre   string
	telefono string
	next     *Cliente
}

var primerCliente *Cliente

func Example() {
	AddCliente("Pepe", "222")
	AddCliente("Juan", "555 555 556")
	AddCliente("Rodrigo", "555 555 557")
	AddCliente("David", "555 555 558")

	ListaClientes()
	// fmt.Printf("El valor de la v es %d \n", v)
	// fmt.Printf("La posicion de memoria de v es %v \n", &v)

	var cliente *Cliente = GetCliente(2)
	(*cliente).nombre = "Albert" // Se actualiza el cliente

	fmt.Printf("\nClientes Update\n")
	ListaClientes()

	DeleteCliente(2)
	fmt.Printf("\nClientes Delete\n")
	ListaClientes()
}

func DeleteCliente(pos int) {
	if primerCliente == nil {
		return
	}

	if pos == 0 {
		if primerCliente.next == nil {
			primerCliente = nil
			return
		}

		primerCliente = primerCliente.next
		return
	}

	var cliente *Cliente = GetCliente(pos)
	var previo *Cliente = GetCliente(pos - 1)
	if cliente == nil || previo == nil {
		return
	}

	previo.next = cliente.next
}

func AddCliente(nombre string, telefono string) {
	var cliente = Cliente{
		nombre:   nombre,
		telefono: telefono,
		next:     nil,
	}

	if primerCliente == nil {
		primerCliente = &cliente

		return
	}

	var nextCliente *Cliente = primerCliente
	for nextCliente.next != nil {
		nextCliente = nextCliente.next
	}

	nextCliente.next = &cliente
}

func ListaClientes() {
	fmt.Println("Lista de clientes...")
	if primerCliente == nil {
		return
	}

	var nextCliente *Cliente = primerCliente
	for nextCliente.next != nil {
		fmt.Printf("Nombre: %s Teléfono: %s \n", nextCliente.nombre, nextCliente.telefono)

		nextCliente = nextCliente.next
	}

	fmt.Printf("Nombre: %s Teléfono: %s \n", nextCliente.nombre, nextCliente.telefono)
}

func GetCliente(pos int) *Cliente {
	if primerCliente == nil {
		return nil
	}

	var nextCliente *Cliente = primerCliente
	i := 0
	for i < pos && nextCliente.next != nil {
		fmt.Println("Posición vale:", pos)
		nextCliente = nextCliente.next
		i++
	}

	if i != pos {
		return nil
	}

	return nextCliente
}
