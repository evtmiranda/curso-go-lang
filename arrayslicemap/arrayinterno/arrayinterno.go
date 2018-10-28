package main

import "fmt"

func main() {
	// é possível ter mais de 1 slice apontando para o mesmo array interno
	s := make([]int, 10, 20)
	fmt.Println(s)
	s1 := append(s, 1, 2, 3)

	s[2] = 9

	fmt.Println(s, s1)
}
