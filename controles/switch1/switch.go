package main

import "fmt"

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(9))
	fmt.Println(notaParaConceito(7))
	fmt.Println(notaParaConceito(2))
}

func notaParaConceito(n float64) string {
	var nota = int(n)

	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	default:
		return "Nota inválida"
	}
}
