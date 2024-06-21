package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(
		variavel3 string =  "blablabla"
		variavel4 string =  "lalala"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	fmt.Println(variavel5, variavel6)
	
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}