package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma variável
	// é a forma mais utilizada
	area := PI * m.Pow(raio, 2)

	// a concatenação de string é feita por "," e não por "+"
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f = true, 3.12

	fmt.Println(e, f)

	g, h := "oi", 5

	fmt.Println(g, h)
}
