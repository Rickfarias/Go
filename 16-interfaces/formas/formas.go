package main

import (
	"fmt"
	"math"
)

func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.area())
}

type triangulo struct {
	altura float64
	largura float64
}

func (t triangulo) area() float64 {
	return (t.altura * t.largura) / 2
}

type retangulo struct {
	altura float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

type forma interface {
	area() float64
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

	t := triangulo{16, 9}
	escreverArea(t)
}