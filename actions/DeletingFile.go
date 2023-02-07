package actions

import (
	"fmt"
	"os"
	"zz/values"
)

// Delete Удаление файла
func Delete(err error, input JsonInput) bool {
	//Проверим, есть ли вобще файл
	Existence(values.FileRenamed, err)
	err = os.Remove(values.FileRenamed)
	if err != nil {
		fmt.Println(err)
		return true
	}

	input.string = "Options.Delete.Condition"
	input.value = 1

	input.JsonWriter()
	return false
}
