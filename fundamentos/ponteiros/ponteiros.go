package main

import "fmt"

func main() {

	a := 5

	var p *int

	var p2 *int

	p = &a
	p2 = &a

	fmt.Println("Endereço de memória para onde o ponteiro aponta =>", p)
	fmt.Println("Valor da variável para onde o ponteiro aponta =>", *p)
	fmt.Println("Endereço de memória do ponteiro =>", &p)

	fmt.Println("Endereço de memória da variável =>", &a)
	fmt.Println("Valor da variável =>", a)

	fmt.Println("Endereço de memória para onde o ponteiro 2 aponta =>", p2)
	fmt.Println("Valor da variável para onde o ponteiro 2 aponta =>", *p2)
	fmt.Println("Endereço de memória do ponteiro 2 =>", &p2)
}
