package clienteInconsistente

import (
	"etlProject/database"
	"etlProject/model/dadoBruto"
	"etlProject/utils"
	"fmt"
	"io/ioutil"
	"os"
)

// Cria um arquivo temporário e envia para ser copiado para o banco
func Create(dadosBrutos []dadoBruto.Model) {
	tmpFile, err := ioutil.TempFile("/tmp/", "etlProject-inconsistentes-")
	utils.CheckErr(err)

	defer os.Remove(tmpFile.Name())

	fmt.Println("Arquivo temporário criado: " + tmpFile.Name())

	for _, dado := range dadosBrutos {
		text := []byte(
			dado.Cpf + ";" +
				dado.Private + ";" +
				dado.Incompleto + ";" +
				dado.DtUltimaCompra + ";" +
				dado.TicketMedio + ";" +
				dado.TicketUltimaCompra + ";" +
				dado.LojaMaisFrequente + ";" +
				dado.LojaUltimaCompra + "\n")
		tmpFile.Write(text)
	}

	tmpFile.Close()

	database.Copy("clientes_inconsistentes", "cpf,private,incompleto,dt_ultima_compra,ticket_medio,ticket_ultima_compra,loja_mais_frequente,loja_ultima_compra", tmpFile.Name())
}
