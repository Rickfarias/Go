package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "Rick",
		"sobrenome": "Farias",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string] string {
		"nome": {
			"Primeiro": "Antonio",
			"Segundo": "Rick",
			"Terceiro": "Farias",
			"Quarto": "Oliveira",
		},
		"curso": {
			"faculdade": "UFC",
			"nome": "Ciências da Computação",
			"campus": "Campus de Crateús",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}