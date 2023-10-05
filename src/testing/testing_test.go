package testing

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	tables := []struct { // Creando tabla de test cases
		a int
		b int
		n int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{25, 26, 51},
	}

	for _, item := range tables { // El guion bajo se utiliza como placeholder para el indice
		total := Sum(item.a, item.b)
		if total != item.n {
			t.Errorf("Sum was incorrect, got: %d, want: %d.", total, item.n)
		}
	}

}

func TestSumStr(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	tables := []struct { // Creando tabla de test cases
		name      string
		a         string
		b         string
		excpected int
	}{
		{"Sum two numbers", "?", "1", 0},
		{"Incorrect a param", "2", "2", 4},
		{"Incorrect b param", "25", "b", 0},
	}

	for _, item := range tables { // El guion bajo se utiliza como placeholder para el indice
		t.Run(item.name, func(t *testing.T) { // Aquí se crea una subprueba. El primer parametro es el nombre de la subprueba
			total := SumStr(item.a, item.b)
			if total != item.excpected {
				t.Errorf("Sum was incorrect, got: %d, want: %d.", total, item.excpected)
			}
		})
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct { // Creando tabla de test cases
		a int
		b int
		n int
	}{
		{2, 1, 2},
		{3, 2, 3},
		{27, 26, 27},
	}

	for _, item := range tables { // El guion bajo se utiliza como placeholder para el indice
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got: %d, want: %d.", max, item.n)
		}
	}
}

func TestFib(t *testing.T) {

}
