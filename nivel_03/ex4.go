package main

import (
	"fmt"
)

const anoAtual = 2023
const anoInicial = 2004

func main() {

	anosVividos := 0

	for {

		if anoInicial+anosVividos > anoAtual {
			break
		}

		fmt.Printf("Ano: %d\n", anoInicial+anosVividos)
		fmt.Printf("Idade : %d\n", anosVividos)
		fmt.Println("")

		anosVividos++
	}

}
