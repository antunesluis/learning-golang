package main

import (
	"fmt"
)

const inicioAscii = 'A'
const nLetra = inicioAscii + 26
const nRepet = 3

func main() {

	for i := inicioAscii; i < nLetra; i++ {
		for j := 0; j < nRepet; j++ {
			fmt.Printf("%c %c %c\n", i, i, i)
		}
		fmt.Println("")
	}
}
