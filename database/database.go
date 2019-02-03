package database

import (
	"database/sql"
	"etlProject/utils"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "postgres"
	PATH        = "./source/"
	FILE        = "create-tables.sql"
)

func Connect() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER,
		DB_PASSWORD,
		DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	utils.CheckErr(err)
	return db
}

// Efetua um COPY do arquivo cujo endereço é passado por parâmetro para a tabela
func Copy(table string, fields string, fileName string) {
	db := Connect()
	defer db.Close()

	copyCommand := "COPY " + table + "(" + fields + ") FROM '" + fileName + "' DELIMITER ';';"

	fmt.Printf("Copiando dados do arquivo... ")
	tx, err := db.Begin()
	stmt, err := tx.Exec(copyCommand)
	utils.CheckErr(err)

	total, err := stmt.RowsAffected()
	utils.CheckErr(err)

	tx.Commit()

	fmt.Println("Concluído(s): " + strconv.FormatInt(total, 10) + " registro(s)!")
}

// Executa o arquivo .sql de criação das tabelas utilizadas
func RecreateTables() {
	fmt.Printf("Recriando as tabelas... ")
	archive, err := ioutil.ReadFile(PATH + FILE)
	utils.CheckErr(err)

	db := Connect()
	defer db.Close()

	requests := strings.Split(string(archive), ";")

	for _, request := range requests {
		_, err := db.Exec(request)
		utils.CheckErr(err)
	}
	fmt.Println("Concluído!")
}
