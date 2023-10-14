package main

import (
	"fmt"
)

func closure() func() int {
	x := 0

	return func() int {
		x += 1
		return x
	}
}

func main() {
	a := closure()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
