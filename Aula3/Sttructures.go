package main

import "fmt"
type Funcionario struct{
	name string
	idade int
	função string
	salario int
}

a := Funcionario {"Fernando", "40", "Engenheiro", "1200" }
b := Funcionario {"Izadora", "22", "Engenheiro", "1700" }
c := Funcionario {"Nicolas", "25", "Engenheiro", "1800" }

fmt.Printf("Funcionario 1 %v\n", a.idade)