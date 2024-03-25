package main

import "fmt"

func main() {
	var var1 int
	var var2 = 2

	fmt.Println("Digite um numero: ")
	fmt.Scanln(&var1)

	resultado := var1 + var2

	fmt.Println("O resultado da soma é:", resultado)
}
