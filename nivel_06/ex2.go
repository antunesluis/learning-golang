package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) mostra_pessoa() {
	fmt.Println(p.nome)
	fmt.Println(p.sobrenome)
	fmt.Println(p.idade)
}

func main() {
	luis := pessoa{"luis", "fernando", 18}
	luis.mostra_pessoa()
}
