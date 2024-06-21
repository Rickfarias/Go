package main

import "fmt"

// 4 tipos de números inteiros na linguagem Go.
// int 8, 16, 32 e 64

func main() {
	var numero int16 = 300
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// Números reais
	// float32, float6
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	// Fim números reais 

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)
	
	str2 := "Texto 2"
	fmt.Println(str2)
	// FIM STRINGS
}

// uint = unsigned int