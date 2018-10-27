package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha")

	fmt.Println(" Nova")
	fmt.Println(" linha")

	x := 3.6978

	//Converte a variável para string para que seja possível escrever no console através de concatenação
	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)

	fmt.Println("O valor de x é", x)

	//Print formatado
	fmt.Printf("O valor de x é %.2f.", x)

	a := 1
	b := 1.999
	c := false
	d := "opa"

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)

	fmt.Printf("\n %v %v %v %v", a, b, c, d)

}
