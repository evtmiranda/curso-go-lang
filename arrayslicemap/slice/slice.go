package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [5]int{1, 2, 3, 4, 5}
	s1 := a1[0:3]

	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	// Slice não é um array! Slice define um pedaço de array
	s2 := a1[:3]
	fmt.Println(a1, s2)

	s3 := s2[:2]
	fmt.Println(s2, s3)

	for i, num := range a1 {
		fmt.Println(i, num)
	}

	fmt.Println("")

	for i, num := range s1 {
		fmt.Println(i, num)
	}
}
