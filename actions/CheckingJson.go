package actions

import "github.com/tidwall/gjson"

const JsonAdr = "./json/Variant.json"

// CheckingJson Проверка содержимого json файла для условий
func CheckingJson(f []byte) (gjson.Result, gjson.Result, gjson.Result, gjson.Result, gjson.Result, gjson.Result) {
	value := gjson.Get(string(f), "Options.Create.Condition")
	value1 := gjson.Get(string(f), "Options.Change_Name.Condition")
	value2 := gjson.Get(string(f), "Options.Delete.Condition")
	value3 := gjson.Get(string(f), "Options.Get_Time.Condition")
	value4 := gjson.Get(string(f), "Options.Input.Condition")
	g := gjson.Get(JsonAdr, "Options.Timestamp")
	return value, value1, value2, value3, value4, g
}
