package general // Le decimos a GO que este archivo pertenece al paquete mypackage

// CarPublic Car con acceso público
// La primera letra de la estructura debe ser mayúscula para que sea pública, si es minúscula será privada
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage Imprime un mensaje
func PrintMessage() {
	println("Hola mundo")
}