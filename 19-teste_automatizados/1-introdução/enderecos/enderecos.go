package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem o tipo válido.
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	// RUA DOS BOBOS
	// ["RUA", "DOS", "BOBOS"]

	enderecoTemUmTipoValido := false
	for _, tipo :=  range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}