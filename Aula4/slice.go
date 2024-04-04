package main

import "fmt"

func main() {
	// Criando uma constante com o valor "hello"
	const mensagem = "hello"

	// Criando um slice de strings
	meuSlice := []string{mensagem, "izadora", "jessica", "Natalia"}

	//Segundo valor a uma variavel
	segundoValor := meuSlice[1]

	// Imprimindo a constante e o slice
	fmt.Println("Esse é o mesmo slice", meuSlice[0], segundoValor)
}
