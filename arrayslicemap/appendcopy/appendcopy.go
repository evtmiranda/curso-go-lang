package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}

	s1 := make([]int, 5)
	s1 = append(s1, 1, 2, 3)

	fmt.Println(a, s1)

	s2 := make([]int, 2)
	copy(s2, s1)

	fmt.Println(s1, s2)
}
