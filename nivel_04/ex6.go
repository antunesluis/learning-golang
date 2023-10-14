package main

import (
	"fmt"
)

func main() {
	var n_pessoas int

	fmt.Println("Digite o numero de pessoas a serem incluidas: ")
	fmt.Scanln(&n_pessoas)

	mapa := make(map[string][]string)

	fmt.Println("Digite o nome e o sobrenome dessas pessoas (um por linha): ")
	for i := 0; i < n_pessoas; i++ {
		var nome, sobrenome string

		fmt.Scanln(&nome)
		fmt.Scanln(&sobrenome)

		mapa[nome] = append(mapa[nome], sobrenome)
	}

	fmt.Println("\nDADOS:")
	for chave, valor := range mapa {
		fmt.Printf("\tNome %s, Sobrenome %s\n", chave, valor)
	}
}
