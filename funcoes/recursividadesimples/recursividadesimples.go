package main

import "fmt"

func main() {
	fmt.Println(fatorial(5))
}

func fatorial(n uint) uint {
	if n == 0 {
		return 1
	}

	fatAnterior := fatorial(n - 1)

	return n * fatAnterior
}
