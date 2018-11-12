package main

import "fmt"

type aluno struct {
	nome      string
	sala      string
	maiorNota float64
	menorNota float64
}

// (a aluno) é a forma de dizer que a função obterNotaMedia é um receiver do tipo aluno
// com isso, automaticamente o método obterNotaMedia fica amarrado ao tipo aluno
func (a aluno) obterNotaMedia() float64 {
	return (a.maiorNota + a.menorNota) / 2
}

func main() {
	aluno1 := aluno{
		nome:      "Everton",
		sala:      "1546",
		maiorNota: 9.87,
		menorNota: 6.75,
	}

	aluno2 := aluno{"Gabriela", "2", 10, 9}

	fmt.Println(aluno1, aluno1.obterNotaMedia())
	fmt.Println(aluno2, aluno2.obterNotaMedia())
}
