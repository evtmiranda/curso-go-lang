package main

import "fmt"

func main() {
	fmt.Println(somar(6, 8))
}

func somar(a, b int) int {
	defer fmt.Println("\nfim...")

	fmt.Printf("Somando %d + %d", a, b)

	return (a + b)
}
