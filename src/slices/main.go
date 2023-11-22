package main

import "fmt"

func main() {
	// Aquí estamos añadiendo un elemento al slice original en una posicion diferente a la ultima y a la primera.
	// Crear un slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// Añadiendo elemento en una posicion especifica de un slice
	nombres := []string{"pedro", "maria", "juan"}
	newName := "jose"
	slides := append(nombres[:3], append([]string{newName}, nombres[3:]...)...) // Adding new slide in 2nd position
	fmt.Println(nombres)
	fmt.Println(slides)
}
