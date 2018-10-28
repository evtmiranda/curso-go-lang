package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1] = "Everton"
	aprovados[2] = "Gabriela"
	aprovados[3] = "Lulu"

	for id, name := range aprovados {
		fmt.Printf("%d) %s\n", id, name)
	}

	fmt.Println(aprovados[1])
	delete(aprovados, 1)
	fmt.Println(aprovados[1])
}
