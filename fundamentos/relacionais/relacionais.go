package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings:", "Banana" == "Banana")

	a := 3
	b := 4

	fmt.Println(">", a > b)
	fmt.Println("<", a < b)
	fmt.Println("==", a == b)
	fmt.Println(">=", a >= b)
	fmt.Println("<=", a <= b)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Everton"}
	p2 := Pessoa{"Gabriela"}

	fmt.Println("Pessoas:", p1 == p2)
}
