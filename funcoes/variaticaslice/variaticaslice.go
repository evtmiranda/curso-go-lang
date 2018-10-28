package main

import "fmt"

func main() {
	titulo := "Lista de Aprovados"
	aprovados := []string{"Everton", "Gabriela", "Lulu"}

	imprimir(titulo, aprovados...)
}

func imprimir(titulo string, valores ...string) {
	fmt.Println(titulo)

	for i, valor := range valores {
		fmt.Printf("%d) %s\n", i+1, valor)
	}
}
