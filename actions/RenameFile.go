package actions

import (
	"fmt"
	"os"
	"zz/values"
)

// Rename Переименование файла
func Rename(err error, input JsonInput) bool {
	//Проверим, есть ли вобще файл
	Existence(values.File, err)
	err = os.Rename(values.File, values.FileRenamed)

	if err != nil {
		fmt.Println(err)
		return true
	}

	input.string = "Options.Change_Name.Condition"
	input.value = 1

	input.JsonWriter()
	return false
}
