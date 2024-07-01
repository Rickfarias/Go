package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	
	p1 := pessoa{"Rick", "Farias", 18, 178}
	fmt.Println(p1)
	
	e1 := estudante{p1, "Ciências da Computação", "UFC"}
	fmt.Println(e1)
}
