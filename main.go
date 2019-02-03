package main

import (
	"etlProject/database"
	"etlProject/file"
	"etlProject/model/clienteInconsistente"
	"etlProject/model/clienteSanitizado"
	"etlProject/model/clienteSemCompra"
	"etlProject/model/dadoBruto"
	"etlProject/utils"
	"etlProject/validator"
	"fmt"
	"strconv"
)

func main() {
	startLog(1)
	// Recria as tabelas para iniciarmos o processo do zero
	database.RecreateTables()
	endLog(1)

	startLog(2)
	// Extrai os dados do arquivo linha a linha
	extracted := file.Read()

	// Copia os dados extraídos para o banco
	dadoBruto.Create(extracted)
	endLog(2)

	startLog(3)
	// Separa os tipos de situações possíveis: sem compras, inconsistentes e não inconsistentes
	fmt.Printf("Analisando registros e verificando o destino dos dados... ")
	tuples := dadoBruto.SelectAll()

	var clientesSemCompras []dadoBruto.Model
	var clientesInconsistentes []dadoBruto.Model
	var clientesSanitizados []dadoBruto.Model

	for _, tuple := range tuples {
		if utils.IsNull(tuple.DtUltimaCompra) && validator.IsValidCpf(tuple.Cpf) {
			clientesSemCompras = append(clientesSemCompras, tuple)
		} else if !validator.IsValidCpf(tuple.Cpf) || !validator.IsValidCnpj(tuple.LojaMaisFrequente) || !validator.IsValidCnpj(tuple.LojaUltimaCompra) {
			clientesInconsistentes = append(clientesInconsistentes, tuple)
		} else {
			clientesSanitizados = append(clientesSanitizados, tuple)
		}
	}
	fmt.Println("Concluído!")
	endLog(3)

	startLog(4)
	// Persiste os registros anteriormente separados em suas bases de dados

	fmt.Println("Persistindo dados de clientes sanitizados:")
	clienteSanitizado.Create(clientesSanitizados)

	fmt.Println("\nPersistindo dados de clientes sem compras:")
	clienteSemCompra.Create(clientesSemCompras)

	fmt.Println("\nPersistindo dados de clientes com inconsistências nas informações:")
	clienteInconsistente.Create(clientesInconsistentes)

	endLog(4)

	fmt.Println("Processo concluído!")
}

func endLog(etapa int) {
	fmt.Println("##### ETAPA " + strconv.Itoa(etapa) + " ##### - FINAL\n")
}

func startLog(etapa int) {
	fmt.Println("##### ETAPA " + strconv.Itoa(etapa) + " ##### - INÍCIO")
}
