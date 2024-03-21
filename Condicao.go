package main

import "fmt"

func main() {
	var idade int
	var permissaoDosPais bool

	fmt.Println("Quantos anos voce tem")
	fmt.Scan(&idade)

	fmt.Println("Voce nao tem permissão dos seus pais?")
	fmt.Scan(&permissaoDosPais)

	podeAssistir := idade >= 18 || (idade >= 13 && permissaoDosPais)
	fmt.Println("Voce pode assistir a este filme:", podeAssistir)
}
