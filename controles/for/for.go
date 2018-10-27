package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // em go não tem while
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Printf("\n")

	for i := 0; i <= 20; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Printf("\n")

	// laço infiníto
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second * 2)
	}

}
