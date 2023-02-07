package actions

// Вносим Timestamp в json
var Ct, Json = SetTimestamp()

type JsonChanger interface {
	JsonWriter() bool
}

// JsonInput Структура данных для изменения json файла
type JsonInput struct {
	jsonfile string
	string   string
	value    int
}

// Input Структура параметров изменения json файла
var Input = JsonInput{
	Json,
	"",
	0,
}
