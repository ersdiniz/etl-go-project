package dadoBruto

import (
	"etlProject/database"
	"etlProject/utils"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Model struct {
	Cpf                string
	Private            string
	Incompleto         string
	DtUltimaCompra     string
	TicketMedio        string
	TicketUltimaCompra string
	LojaMaisFrequente  string
	LojaUltimaCompra   string
}

// Cria um arquivo temporário e envia para ser copiado para o banco
func Create(lines []string) {
	tmpFile, err := ioutil.TempFile("/tmp/", "etlProject-")
	utils.CheckErr(err)

	defer os.Remove(tmpFile.Name())

	fmt.Println("Arquivo temporário criado: " + tmpFile.Name())

	for _, line := range lines {
		lineArray := strings.Fields(line)
		text := []byte(strings.Join(lineArray, ";") + "\n")
		tmpFile.Write(text)
	}

	tmpFile.Close()

	database.Copy("dados_brutos", "cpf,private,incompleto,dt_ultima_compra,ticket_medio,ticket_ultima_compra,loja_mais_frequente,loja_ultima_compra", tmpFile.Name())
}

func SelectAll() []Model {
	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT cpf,private,incompleto,dt_ultima_compra,ticket_medio,ticket_ultima_compra,loja_mais_frequente,loja_ultima_compra FROM dados_brutos")
	utils.CheckErr(err)

	var tuples []Model
	for rows.Next() {
		var model Model

		err = rows.Scan(&model.Cpf, &model.Private, &model.Incompleto, &model.DtUltimaCompra, &model.TicketMedio, &model.TicketUltimaCompra, &model.LojaMaisFrequente, &model.LojaUltimaCompra)
		utils.CheckErr(err)
		tuples = append(tuples, model)
	}

	rows.Close()

	return tuples
}
