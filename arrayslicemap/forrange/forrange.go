package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta

	// parecido com o foreach, porém com acesso ao índice
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	// parecido com o foreach
	for _, num := range numeros {
		fmt.Println(num)
	}
}
