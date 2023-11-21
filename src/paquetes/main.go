package main

import (
	"fmt"
	"leeme_go/src/paquetes/custom_package"
)

func main() {
	// No importa que el método GetTitle() este en otro archivo, lo podemos usar igual ya que esta en el mismo paquete.
	curso := custom_package.Curso{Titulo: "Curso de Go"}
	fmt.Println(curso.GetTitle())
}
