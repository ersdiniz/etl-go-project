package clienteSemCompra

import (
	"etl-go-project/database"
	"etl-go-project/file"
	"etl-go-project/model/dadoBruto"
	"etl-go-project/sanitize"
	"etl-go-project/utils"
	"fmt"
	"io/ioutil"
	"os"
)

// Cria um arquivo temporário e envia para ser inserido no banco
func Create(dadosBrutos []dadoBruto.Model) {
	tmpFile, err := ioutil.TempFile(file.CreateFolderIfNotExists("/tmp/"), "etlProject-sem_compra-")
	utils.CheckErr(err)

	defer os.Remove(tmpFile.Name())

	fmt.Println("Arquivo temporário criado: " + tmpFile.Name())

	for _, dado := range dadosBrutos {
		sanitize.CleanNumeric(&dado.Cpf)

		text := []byte(
			"'" + dado.Cpf + "','" +
				dado.Private + "','" +
				dado.Incompleto + "'\n")
		tmpFile.Write(text)
	}

	tmpFile.Close()

	database.Insert("clientes_sem_compras", "cpf,private,incompleto", tmpFile.Name())
}
