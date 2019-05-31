package sanitize

import (
	"etl-go-project/utils"
	"regexp"
)

// Remove os caracteres não numéricos da string
func CleanNumeric(value *string) {
	reg, err := regexp.Compile("[^0-9]+")
	utils.CheckErr(err)
	*value = reg.ReplaceAllString(*value, "")
}

// Remove os caracteres não numéricos da string e substitui "," por "."
func CleanFloat(value *string) {
	reg, err := regexp.Compile("[,]")
	utils.CheckErr(err)
	*value = reg.ReplaceAllString(*value, ".")

	reg, err = regexp.Compile("[^0-9.]+")
	utils.CheckErr(err)
	*value = reg.ReplaceAllString(*value, "")
}
