package main

import "fmt"

func main() {
	var name string
	var idade int

	fmt.Println("Digite seu nome: ")
	fmt.Scanln(&name)

	fmt.Println("Digite sua idade: ")
	fmt.Scanln(&idade)

	fmt.Println("Meu nome é,", name, "e tenho", idade, "anos")
}
