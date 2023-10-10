package main

import (
	"fmt"
)

func main() {
	var n_pessoas int

	fmt.Println("Digite o numero de pessoas a serem incluidas: ")
	fmt.Scanln(&n_pessoas)

	matriz := make([][]string, n_pessoas)

	fmt.Println("Digite o nome, sobrenome e hobby dessas pessoas (um por linha): ")
	for i := 0; i < n_pessoas; i++ {
		var nome, sobrenome, hobby string

		fmt.Scanln(&nome)
		fmt.Scanln(&sobrenome)
		fmt.Scanln(&hobby)

		detalhes := []string{nome, sobrenome, hobby}
		matriz[i] = detalhes
	}

	fmt.Println("\nMatriz de detalhes das pessoas:")
	for i := 0; i < n_pessoas; i++ {
		fmt.Printf("\tPessoa %d: %s, %s, %s\n", i+1, matriz[i][0], matriz[i][1], matriz[i][2])
	}
}
