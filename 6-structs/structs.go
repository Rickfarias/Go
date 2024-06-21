package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Rick"
	u.idade = 18
	fmt.Println(u)

	usuario2 := usuario{"Guilherme", 21}
	fmt.Println(usuario2)

	usuario3 := usuario{nome:"João", idade: 23}
	fmt.Println(usuario3)
}