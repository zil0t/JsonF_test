package actions

import (
	"fmt"
	"os"
)

const FileRenamed = "awesome.txt"

// Rename Переименование файла
func Rename(err error, json string) bool {
	//Проверим, есть ли вобще файл
	Existence(file, err)
	err = os.Rename(file, FileRenamed)

	if err != nil {
		fmt.Println(err)
		return true
	}
	//Вносим изменения в json
	input := JsonInput{
		json,
		"Options.Change_Name.Condition",
		1,
	}

	input.JsonWriter()
	return false
}
