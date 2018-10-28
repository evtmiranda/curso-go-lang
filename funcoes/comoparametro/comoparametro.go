package main

import "fmt"

// casos de uso deste tipo de programação: reduce, filter e map

func main() {
	fmt.Println(exec(multiplicacao, 3, 5))
}

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}
