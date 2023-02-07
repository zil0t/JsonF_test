package actions

import (
	"fmt"
	"log"
	"os"
	"time"
	"zz/values"
)

// Creating Создание файла
func Creating(input JsonInput) {
	f, err := os.OpenFile(values.File,
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

	input.string = "Options.Create.Condition"
	input.value = 1
	input.JsonWriter()
}
