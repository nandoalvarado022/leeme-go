// Cuando el test retorna go: cannot find module, but found... es porque no se ha hecho el go mod init
// go test -coverprofile=coverage.out // Para generar el archivo de coverage
// go tool cover -func=coverage.out // Para ver el coverage en la terminal. Muestra el porcentaje de coverage por cada funcion
// go tool cover -html=coverage.out // Para ver el coverage en el navegador
// go test -cpu=1,2,4,8 // Para correr el test en diferentes procesadores
// go test -cpuprofile=cpu.out // Para generar el archivo de cpu profile
// table drive test Se crea una tabla tipo arreglo donde estan los posibles escenarios de prueba
// Subpruebas: Se pueden crear subpruebas dentro de un test. Esto es util cuando se quiere probar un metodo que tiene muchos casos de uso

package testing

import (
	"strconv"
)

func Sum(x, y int) int {
	return x + y
}

func SumStr(x, y string) int {
	c, err := strconv.Atoi(x)
	if err != nil {
		return 0
	}

	d, err := strconv.Atoi(y)
	if err != nil {
		return 0
	}

	return c + d
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
