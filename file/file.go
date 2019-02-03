package file

import (
	"bufio"
	"etlProject/utils"
	"os"
)

const (
	PATH = "./source/"
	FILE = "base_teste.txt"
)

// Lê o arquivo do caminho especificado e retorna um array com as linhas
//
// Em casos reais, este caminho fixo não é aplicável, já que a origem dos
// dados poderia ser de várias formas, mas dificilmente seria de um arquivo
// anexado ao código fonte da aplicação
func ReadSourceFile() []string {
	archive, err := os.Open(PATH + FILE)
	utils.CheckErr(err)
	defer archive.Close()

	scanner := bufio.NewScanner(archive)
	scanner.Split(bufio.ScanLines)

	var lines []string

	isHeaderLine := true
	for scanner.Scan() {
		if !isHeaderLine {
			lines = append(lines, scanner.Text())
		} else {
			isHeaderLine = false
		}
	}

	return lines
}

func CreateFolderIfNotExists(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir)
	}
	return path
}