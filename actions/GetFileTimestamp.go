package actions

import (
	"github.com/tidwall/gjson"
	"log"
)

// Timestamp Получение даты из файла
func Timestamp(json string) {
	g := gjson.Get(JsonAdr, "Options.Timestamp")
	log.Println(g)
	//Вносим изменения в json
	input := JsonInput{
		json,
		"Options.Get_Time.Condition",
		1,
	}

	input.JsonWriter()
}
