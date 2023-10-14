package main

import (
	"fmt"
)

var esporte string

func main() {

	fmt.Println("Digite o seu esporte favorito dentre as opções:")
	fmt.Printf("\tfutebol\n\tbasquete\n\tvolei,\n\tqueimada\n")
	fmt.Scanln(&esporte)

	switch esporte {
	case "futebol":
		fmt.Println("Eu também gosto de futebol :)")
	case "basquete":
		fmt.Println("Eu também gosto de basquete :)")
	case "volei":
		fmt.Println("Eu também gosto de volei :)")
	case "queimada":
		fmt.Println("Eu também gosto de queimada :)")
	default:
		fmt.Println("Opcão Inválida")
	}
}
