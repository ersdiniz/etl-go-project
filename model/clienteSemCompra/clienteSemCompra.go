package clienteSemCompra

import (
	"etlProject/database"
	"etlProject/model/dadoBruto"
	"etlProject/sanitize"
	"etlProject/utils"
	"fmt"
	"io/ioutil"
	"os"
)

// Cria um arquivo temporário e envia para ser copiado para o banco
func Create(dadosBrutos []dadoBruto.Model) {
	tmpFile, err := ioutil.TempFile("/tmp/", "etlProject-sem_compra-")
	utils.CheckErr(err)

	defer os.Remove(tmpFile.Name())

	fmt.Println("Arquivo temporário criado: " + tmpFile.Name())

	for _, dado := range dadosBrutos {
		sanitize.CleanNumeric(&dadoBruto.Cpf)

		text := []byte(
			dado.Cpf + ";" +
				dado.Private + ";" +
				dado.Incompleto + "\n")
		tmpFile.Write(text)
	}

	tmpFile.Close()

	database.Copy("clientes_sem_compras", "cpf,private,incompleto", tmpFile.Name())
}
