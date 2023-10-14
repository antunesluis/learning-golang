package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	quatroRodas bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	carro_1 := sedan{
		veiculo{4, "vinho"},
		true,
	}

	carro_2 := caminhonete{
		veiculo{6, "vinho"},
		false,
	}

	fmt.Println(carro_1)
	fmt.Println(carro_2)
}
