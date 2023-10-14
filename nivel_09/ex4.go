package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "diz olá")
}

type humano interface {
	falar()
}

func falarAlgo(h humano) {
	h.falar()
}

func main() {
	pessoa1 := pessoa{"jorge", 32}

	pessoa1.falar()
	falarAlgo(&pessoa1)
}
