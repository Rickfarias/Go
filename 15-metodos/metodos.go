package main

import "fmt"

type usuario struct {
	nome string 
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário: %s, no banco de dados \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{nome: "Usuário", idade: 22}
	usuario1.salvar()
	
	usuario2 := usuario{"Davi", 20}
	usuario2.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Printf("é %t", maiorDeIdade)
}

