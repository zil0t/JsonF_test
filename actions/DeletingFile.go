package actions

import (
	"fmt"
	"os"
)

// Delete Удаление файла
func Delete(err error, json string) bool {
	//Проверим, есть ли вобще файл
	Existence(FileRenamed, err)
	err = os.Remove(FileRenamed)
	if err != nil {
		fmt.Println(err)
		return true
	}
	//Вносим изменения в json
	input := JsonInput{
		json,
		"Options.Delete.Condition",
		1,
	}

	input.JsonWriter()
	return false
}
