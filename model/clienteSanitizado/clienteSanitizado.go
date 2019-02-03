package clienteSanitizado

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
	tmpFile, err := ioutil.TempFile("/tmp/", "etlProject-sanitizados-")
	utils.CheckErr(err)

	defer os.Remove(tmpFile.Name())

	fmt.Println("Arquivo temporário criado: " + tmpFile.Name())

	for _, dado := range dadosBrutos {
		sanitize.CleanNumeric(&dado.Cpf)
		sanitize.CleanNumeric(&dado.LojaMaisFrequente)
		sanitize.CleanNumeric(&dado.LojaUltimaCompra)

		sanitize.CleanFloat(&dado.TicketMedio)
		sanitize.CleanFloat(&dado.TicketUltimaCompra)

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

	database.Copy("clientes_sanitizados", "cpf,private,incompleto,dt_ultima_compra,ticket_medio,ticket_ultima_compra,loja_mais_frequente,loja_ultima_compra", tmpFile.Name())
}
