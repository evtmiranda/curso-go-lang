package main

import "fmt"

func main() {
	x, y := trocar(4, 9)
	fmt.Println(x, y)

	primeiro, segundo := trocar(10, 8)
	fmt.Println(primeiro, segundo)
}

func trocar(x, y int) (primeiro, segundo int) {
	primeiro = y
	segundo = x

	return //retorno limpo
}
