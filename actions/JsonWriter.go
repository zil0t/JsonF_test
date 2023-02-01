package actions

import (
	"github.com/tidwall/sjson"
	"log"
	"os"
)

// JsonChanger Интерфейс изменения json файла
type JsonChanger interface {
	JsonWriter() bool
}

// JsonInput Структура данных для изменения json файла
type JsonInput struct {
	jsonfile string
	string   string
	value    int
}

// JsonWriter Внесение изменений в json
func (input JsonInput) JsonWriter() bool {
	n, err := sjson.Set(input.jsonfile, input.string, input.value)
	if err != nil {
		log.Println("Проверьте аргументы записи в json")
		return false
	}
	z, err := os.OpenFile(JsonAdr, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	_, err = z.WriteString(n)
	return true
}
