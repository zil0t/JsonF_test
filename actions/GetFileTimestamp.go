package actions

import (
	"github.com/tidwall/gjson"
	"log"
	"zz/values"
)

// Timestamp Получение даты из файла
func Timestamp(input JsonInput) {
	g := gjson.Get(values.JsonAdr, "Options.Timestamp")
	log.Println(g)

	input.string = "Options.Get_Time.Condition"
	input.value = 1
	input.JsonWriter()
}
