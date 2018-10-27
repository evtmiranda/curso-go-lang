package main

import "fmt"

func main() {
	fmt.Println(tipo(7))
	fmt.Println(tipo(7.5))
	fmt.Println(tipo("7"))
	fmt.Println(tipo(func() {}))
	fmt.Println(1 == 1)
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "i dont know"
	}
}
