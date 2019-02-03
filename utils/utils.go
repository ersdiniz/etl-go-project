package utils

// Método padrão para tratamento de erros
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func IsNull(value string) bool {
	return len(value) == 0 || value == "NULL"
}
