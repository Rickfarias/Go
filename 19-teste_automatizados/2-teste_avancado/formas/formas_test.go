package formas

import (
	"math"
	"testing"
)

// TDD - Test Driven Development

func TestArea(t *testing.T) {
	t.Run(
		"Retângulo",
		func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	},
	)

	t.Run("Círculo", func(t *testing.T) {
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Triângulo", func(t *testing.T) {
		tri := Triangulo{10, 12}

		areaEsperada := float64(60)
		areaRecebida := tri.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da área esperada %f", areaRecebida, areaEsperada)
		}
	})
}