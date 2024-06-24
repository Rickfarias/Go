package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func main() {
	totalDaSoma := soma(1, 23, 43, 35, 354, 46, 25, 366, 625)
	fmt.Println(totalDaSoma)
}