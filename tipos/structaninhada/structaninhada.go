package main

import "fmt"

type aluno struct {
	nome  string
	sala  string
	notas []nota
}

type nota struct {
	alunoID uint
	nota    float64
}

func (a aluno) maiorNota() float64 {
	maiorNota := float64(0)

	for _, nota := range a.notas {
		if nota.nota > maiorNota {
			maiorNota = nota.nota
		}
	}

	return maiorNota
}

func (a aluno) notaTotal() float64 {
	total := float64(0)

	for _, nota := range a.notas {
		total += nota.nota
	}

	return total
}

func (a aluno) notaMedia() float64 {
	quantidadeNotas := float64(len(a.notas))

	return a.notaTotal() / quantidadeNotas
}

func main() {

	everton := aluno{
		nome: "Everton",
		sala: "1",
		notas: []nota{
			nota{alunoID: 1, nota: 7},
			nota{alunoID: 1, nota: 4},
			nota{alunoID: 1, nota: 7},
			nota{alunoID: 1, nota: 8},
			nota{alunoID: 1, nota: 9.76},
		},
	}

	fmt.Printf("Maior nota: %.2f\n", everton.maiorNota())
	fmt.Printf("Nota total: %.2f\n", everton.notaTotal())
	fmt.Printf("Nota média: %.2f", everton.notaMedia())
}
