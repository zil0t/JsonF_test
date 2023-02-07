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
	"zz/values"
)

func main() {

	//Читаем содержимое json файла
	f, err := os.ReadFile(values.JsonAdr)
	if err != nil {
		panic(err)
	}

	//Проверка содержимого json файла для условий
	value, value1, value2, value3, value4, g := actions.CheckingJson(f)

	//Создаём файл
	if value.Num+value1.Num+value2.Num+value3.Num+value4.Num == 0 {
		actions.Creating(actions.Input)
	}

	//Переименовываем файл
	if value.Num == 1 {
		actions.Rename(err, actions.Input)
	}

	//Удаляем файл
	if value1.Num == 1 {
		actions.Delete(err, actions.Input)

	}

	//Выводим время создания
	if value2.Num == 1 {
		actions.Timestamp(actions.Input)
	}

	//Проверяем время и переписываем json
	if value3.Num == 1 && g.Str < actions.Ct {
		actions.CheckTime(actions.Input)
	}
}
