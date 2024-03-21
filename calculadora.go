package main

import "fmt"

func main() {
	var num int
	var num2 int
	var ope string

	fmt.Println("Digite um numero inteiro: ")
	fmt.Scanln(&num)

	fmt.Println("Digite o segundo numero inteiro: ")
	fmt.Scanln(&num2)

	fmt.Println("Digite o sinal da operação(+,-,*,/): ")
	fmt.Scanln(&ope)

	if ope == "+" {
		soma := num + num2
		fmt.Println("O resultado da soma é:", soma)
	} else if ope == "-" {
		soma := num - num2
		fmt.Println("O resultado da subtração é:", soma)
	} else if ope == "-" {
		soma := num - num2
		fmt.Println("O resultado da subtração é:", soma)
	} else if ope == "*" {
		soma := num * num2
		fmt.Println("O resultado da multiplicação é:", soma)
	} else if ope == "/" {
		soma := num / num2
		fmt.Println("O resultado da divisão é:", soma)
	} else {
		fmt.Println("Operação Invalida")
	}

}
