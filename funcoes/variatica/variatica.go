package main

import "fmt"

func main() {
	fmt.Printf("Média => %.2f\n", media(1.4, 9.8, 5.77, 8.34, 3.56))
	fmt.Println("Soma =>", soma(1, 2, 3, 5, 8))
}

func media(numeros ...float64) (media float64) {
	total := 0.0

	for _, numero := range numeros {
		total += numero
	}

	media = total / float64(len(numeros))

	return
}

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
