package main

import (
	"fmt"
)

type pessoa struct {
	nome              string
	sobrenome         string
	sabor_fav_sorvete string
}

func main() {

	var pessoa_1 pessoa

	fmt.Println("Digite seu nome:")
	fmt.Scanln(&pessoa_1.nome)

	fmt.Println("Digite seu sobrenome:")
	fmt.Scanln(&pessoa_1.sobrenome)

	fmt.Println("Digite seu sabor de sorvete favorito:")
	fmt.Scanln(&pessoa_1.sabor_fav_sorvete)

	fmt.Printf("Nome: %s\n", pessoa_1.nome)
	fmt.Printf("Sobrenome: %s\n", pessoa_1.sobrenome)
	fmt.Printf("Sabor Favorito de Sorvete: %s\n", pessoa_1.sabor_fav_sorvete)
}
