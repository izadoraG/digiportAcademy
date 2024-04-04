package main

import "fmt"

type Funcionario struct {
	nome    string
	idade   int
	funcao  string
	salario float32
}

func main() {

	var Funcionarios [3]Funcionario

	a := Funcionario{"Izadora", 22, "Engenheira", 10000}

	b := Funcionario{"Jessica", 25, "Engenheira", 15000}

	c := Funcionario{"Natalia", 28, "Engenheira", 20000}

	fmt.Printf("Funcionario 1 %v\n", a)

	Funcionarios[0] = a
	Funcionarios[1] = b
	Funcionarios[2] = c

}
