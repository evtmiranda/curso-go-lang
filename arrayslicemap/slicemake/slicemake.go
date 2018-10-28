package main

import "fmt"

func main() {
	// cria um slice e um array interno que será referenciado e utilizado pelo slice
	s := make([]int, 10)
	s[9] = 12

	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	// cria um slice e define o tamanho do array interno
	s1 := make([]int, 10, 20)

	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))

	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println(s1)
}
