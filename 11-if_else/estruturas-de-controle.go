package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero >= 15 {
		fmt.Println("Maior ou igual a 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := 0; outroNumero > +0 {
		fmt.Println("Número é maior que 0")
	} else {
		fmt.Println("Número é maior ou igual a 0")
	}
}
