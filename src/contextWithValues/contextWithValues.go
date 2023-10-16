// Ejemplo de como pasar valores entre funciones usando context.
// https://www.youtube.com/watch?v=AlwC08LpsWA
// Colocamos env := "live" en la linea 26 ya que se debe manear un valor por defecto en caso de que no se encuentre el valor en el contexto.
// Pasar parametros entre funciones utilizando el contexto no es buena practica cuando son parametros realmente necesarios para la ejecución de la función.

package main

import (
	"context"
	"fmt"
)

func main() {
	autenticar(context.Background())
}
func autenticar(ctx context.Context) {
	ctx = context.WithValue(ctx, "env", "sanbox")
	func1(ctx)
}

func func1(ctx context.Context) {
	fmt.Println("func1")
	func2(ctx)
}
func func2(ctx context.Context) {
	fmt.Println("func2")

	env := "live"
	if v := ctx.Value("env"); v != nil { // Si el valor de "env" no es nulo, entonces lo asignamos a la variable env
		envString, ok := v.(string)
		if ok {
			env = envString
		}
	}
	fmt.Println("env:", env)
}
