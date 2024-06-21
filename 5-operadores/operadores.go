package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero int16 = 10
	var numero2 int16 = 25
	soma2 := numero + numero2

	fmt.Println(soma2)
	// FIM DOS ARITMETICOS

	// ATRIBUIÇÂO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	// OPERADORES RELACIONAIS
	fmt.Println(2 > 1)
	fmt.Println(2 >= 1)
	fmt.Println(2 < 1)
	fmt.Println(2 <= 1)
	fmt.Println(2 == 1)
	fmt.Println(2 != 1)
	// FIM DOS RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println(false && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// FIM DOS OPERADORES LOGICOS

	// OPERADORES UNÁRIOS
	fmt.Println("-------------")
	numero6 := 20
	numero6++
	numero6 += 15
	fmt.Println(numero6)

	numero6 *= 2
	fmt.Println(numero6)
}