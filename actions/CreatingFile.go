package actions

import (
	"fmt"
	"log"
	"os"
	"time"
)

const file = "log.txt"

// Creating Создание файла
func Creating(json string) {
	f, err := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//Записываем произвольную строку
	timestamp := time.Now()
	_, err = fmt.Fprint(f, timestamp)
	if err != nil {
		log.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	//Меняем состояние json
	input := JsonInput{
		json,
		"Options.Create.Condition",
		1,
	}
	input.JsonWriter()
}
