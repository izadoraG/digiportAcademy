package main

import "fmt"

func main() {
	var num int

	fmt.Println("Digite um numero inteiro: ")
	fmt.Scanln(&num)

	if num > 0 {
		fmt.Println("Positivo!")
	} else if num == 0 {
		fmt.Println("Zero!")
	} else {
		fmt.Println("Negativo")
	}

}
