package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// a variável i só existirá dentro do if/else
	if i := numeroAleatorio(); i > 5 {
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}
}

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}
