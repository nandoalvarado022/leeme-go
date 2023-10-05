package stringers

import (
	"fmt"
	"leeme_go/src/pointers"
)

func Basic() {
	// %s para strings
	// %d para enteros
	// %v para cualquier tipo de dato
	// %f para floats
	// %t para booleanos
	// %q para strings con comillas dobles
	// %p para punteros
	// %T para saber el tipo de dato
	// %b para binarios
	myPc := pointers.Pc{Ram: 16, Disk: 200, Brand: "MSI"}
	customString := fmt.Sprintf("Tengo %d GB de RAM, %d GB de disco y es una %s", myPc.Ram, myPc.Disk, myPc.Brand)
	fmt.Println(customString)
}
