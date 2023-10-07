package main

import (
	"fmt"
)

const (
	_ = iota + 2023
	a
	b
	c
)

func main() {

	fmt.Printf("Próximos anos: %d, %d, %d", a, b, c)
}
