package main

import "fmt"

func main() {
	nomesESalarios := map[string]float64{
		"Everton Miranda":   5000.00,
		"Gabriela Ferreira": 2900.99,
	}

	nomesESalarios["Lulu da Pomerânia"] = 0.00

	delete(nomesESalarios, "Nenhum")

	for nome, salario := range nomesESalarios {
		fmt.Printf("%s %.2f\n", nome, salario)
	}
}
