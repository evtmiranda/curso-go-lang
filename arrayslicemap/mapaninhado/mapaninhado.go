package main

import "fmt"

func main() {
	s := map[string]map[string]float64{
		"E": {
			"Everton Miranda": 6500.99,
		},
		"G": {
			"Gabriela Ferreira": 94568.78,
		},
		"L": {
			"Lulu": 7674.87,
		},
	}

	fmt.Println(s)

	delete(s, "L")

	fmt.Println(s)

	for letra, mapValues := range s {
		fmt.Println(letra, mapValues)

		for nome, salario := range mapValues {
			fmt.Println(nome, salario)
		}
	}
}
