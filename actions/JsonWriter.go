package actions

import (
	"github.com/tidwall/sjson"
	"log"
	"os"
	"zz/values"
)

// JsonWriter Внесение изменений в json
func (input JsonInput) JsonWriter() bool {
	n, err := sjson.Set(input.jsonfile, input.string, input.value)
	if err != nil {
		log.Println("Проверьте аргументы записи в json")
		return false
	}
	z, err := os.OpenFile(values.JsonAdr, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	_, err = z.WriteString(n)
	return true
}
