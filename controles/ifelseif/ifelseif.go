package main

import "fmt"

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(3))
	fmt.Println(notaParaConceito(7.6))
}

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n < 10:
		return "A"
	case n > 6 && n < 9:
		return "B"
	default:
		return "C"
	}
}
