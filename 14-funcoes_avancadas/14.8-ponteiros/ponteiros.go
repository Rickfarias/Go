package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	var input int

	fmt.Print("Digite o número que deseja inverter: ")
	fmt.Scan(&input)

	numeroInvertido := inverterSinal(input)
	fmt.Println("O número invertido é:", numeroInvertido)

	novoNumero := 40
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println("o novo número invertido é:", novoNumero)
}