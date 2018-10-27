package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// inteiros sem sinal: uint8, uint16, uint32, uint64
	ui1 := math.MaxUint8
	ui2 := math.MaxUint16
	ui3 := math.MaxUint32

	fmt.Printf("%d(%s) %d(%s) %d(%s)", ui1, reflect.TypeOf(ui1), ui2, reflect.TypeOf(ui2), ui3, reflect.TypeOf(ui3))

	//byte = uint32
	var b byte = 255

	fmt.Println("\n", b)

	var run rune = 'a' //representa um mapeamento da tabela Unicode (int32)

	fmt.Println(run)
	fmt.Println(reflect.TypeOf(run))

	i5 := math.MaxInt64

	fmt.Println(" \nO valor máximo do int é", i5)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i5))

	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo de x é", reflect.TypeOf(19.99))

	bo := false
	fmt.Println("O tipo de bo é ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	var s1 = "Olá meu nome é Everton"
	s2 := "Olá meu apelido é Evt"

	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string de apelido é", len(s2))

	// string com múltiplas linhas
	s3 := `Olá
	Meu
	Nome
	É
	Everton`

	fmt.Println(s3)
	fmt.Println("O tamanho da string com múltiplas linhas é", len(s3))
}
