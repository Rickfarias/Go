package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste {
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada ABC", "Estrada"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"RODOVIA CASTELO BRANCO", "Rodovia"},
		{"ESTRADA QUALQUER", "Estrada"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", retornoRecebido,
			cenario.retornoEsperado)
		}
	} 
}
