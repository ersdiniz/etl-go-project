package validator

import "github.com/Nhanderu/brdoc"

// Valida a estrutura da string e verifica se é um CNPJ válido
func IsValidCnpj(value string) bool {
	return brdoc.IsCNPJ(value)
}

// Valida a estrutura da string e verifica se é um CPF válido
func IsValidCpf(value string) bool {
	return brdoc.IsCPF(value)
}
