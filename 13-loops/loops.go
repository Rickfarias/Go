package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	fmt.Println("Início")
	time.Sleep(time.Second)
	for i < 10 {
		i++
		fmt.Println("Incrementando I", i)
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second)
	fmt.Println("Final")
	
	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando j",  j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}
}