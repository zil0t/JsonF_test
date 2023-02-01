//По циклу выполняет действия согласно json
// 1. Создание файла с записью "произвольной строки"
// 2. Переименовывает файл
// 3. Удаляет
// 4. Выводит время создания
// 5. Перезаписывает с условием

package main

import (
	"os"
	"zz/actions"
)

const JsonAdr = "./json/Variant.json"

func main() {

	//Вносим Timestamp в json
	ct, json := actions.SetTimestamp()

	f, err := os.ReadFile(JsonAdr)
	if err != nil {
		panic(err)
	}
	//Проверка содержимого json файла для условий
	value, value1, value2, value3, value4, g := actions.CheckingJson(f)

	//Создаём файл
	if value.Num+value1.Num+value2.Num+value3.Num+value4.Num == 0 {
		actions.Creating(json)
	}

	//Переименовываем файл
	if value.Num == 1 {
		actions.Rename(err, json)
	}

	//Удаляем файл
	if value1.Num == 1 {
		actions.Delete(err, json)

	}

	//Выводим время создания
	if value2.Num == 1 {
		actions.Timestamp(json)
	}

	//Проверяем время и переписываем json
	if value3.Num == 1 && g.Str < ct {
		actions.CheckTime(json)
	}
}
