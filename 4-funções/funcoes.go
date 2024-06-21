package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2

	return soma, subtracao, multiplicacao, divisao
}

func main() {
	soma := somar(5, 12)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função f")
	}

	f()

	resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao)
}