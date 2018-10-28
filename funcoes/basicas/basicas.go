package main

import "fmt"

// função pura: não altera nada externo e sempre retornará o mesmo valor dado um determinado parâmetro

func main() {
	f1()
	f2("Jesus", "Te ama")
	fmt.Println(f3())
	fmt.Println(f4("God", "Blessed"))
	primeiro, segundo := f5()

	fmt.Println(primeiro, segundo)
}

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Primeiro", "Segundo"
}
