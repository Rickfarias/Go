package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 5
	fmt.Println(array1)

	fmt.Println("----------")

	array2 := []string{"Posição 1,", "Posição 2,", "Posição 3,", "Posição 4,", "Posição 5,", "Posição 6"}
	fmt.Println(array2)
	
	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array3)

	fmt.Println("--------")
	
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	
	slice = append(slice, 18, 19, 19, 20, 21, 22)
	fmt.Println(slice)

	// ARRAYS INTERNOS
	fmt.Println("---------")

	slice3 := make([]int, 10, 15)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	fmt.Println(slice3)

	slice3 = append(slice3, 10, 15, 20, 25, 30, 35)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// TIP: Quando a capacidade do slice não é especificada,
	// será do mesmo tamanho da quantidade de posições passadas para o slice
}