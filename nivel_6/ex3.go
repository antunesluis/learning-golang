package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	tamanho_lados float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.tamanho_lados * q.tamanho_lados
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

type info interface {
	area() float64
}

func get_info(i info) {
	fmt.Println(i.area())
}

func main() {
	quadrado_1 := quadrado{
		tamanho_lados: 10.5,
	}

	circulo_1 := circulo{
		raio: 6.43,
	}

	fmt.Printf("\nArea do quadrado: %f", quadrado_1.area())
	fmt.Printf("\nArea do circulo: %f", circulo_1.area())

	fmt.Println("\nUsando interface:")
	get_info(quadrado_1)
	get_info(circulo_1)
}
