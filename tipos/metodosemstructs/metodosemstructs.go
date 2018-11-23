package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) getFullName() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) setName(fullName string) {
	names := strings.Split(fullName, " ")
	p.nome = names[0]
	p.sobrenome = names[1]
}

func main() {
	everton := pessoa{
		nome:      "Everton",
		sobrenome: "Miranda",
	}

	fmt.Println(everton.getFullName())

	everton.setName("Gabriela Ferreira")

	fmt.Println(everton.getFullName())
}
