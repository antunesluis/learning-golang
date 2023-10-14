package main

import "fmt"

type erroEspecial struct {
	qualquercoisa string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("deu ruim: %v", e.qualquercoisa)
}

func erroComoParametro(e error) {
	fmt.Println("Passaram como argumento pra mim, o seguinte: ", e.(erroEspecial).qualquercoisa, "\nno método Error, eu tenho:", e)
}

func main() {
	arg := erroEspecial{"erro triste"}
	erroComoParametro(arg)
}
