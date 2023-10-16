package incrementValor

import "fmt"

func Main() {
	var v int = 4
	var p *int

	p = &v
	*p = 8

	incReferencia(&v)

	fmt.Println("Valor de V es %d", v)
	fmt.Println("Posicion de memoria de V es %v", &v)

	fmt.Println("Valor de P es %d", *p)
	fmt.Println("Posicion de memoria de P es %v", p)
}

func incValor(v int) {
	v++
	fmt.Printf("Valor incrementado es %d\n", v)
}

func incReferencia(v *int) {
	*v++
	fmt.Printf("Valor incrementado por referencia es %d\n", *v)
}
