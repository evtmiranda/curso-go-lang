package main

import "fmt"

func main() {
	fmt.Println(fatorial(5))
}

func fatorial(n int) (int, error) {
	if n < 0 {
		return -1, fmt.Errorf("o número é invalido")
	}

	if n == 0 {
		return 1, nil
	}

	fatAnterior, _ := fatorial(n - 1)

	return n * fatAnterior, nil
}
