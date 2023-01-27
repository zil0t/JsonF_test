//По циклу выполняет действия согласно json
// 1. Создание файла с записью "произвольной строки"
// 2. Переименовывает файл
// 3. Удаляет
// 4. Выводит время создания
// 5. Перезаписывает с условием

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const jsonAdr = "./json/Variant.json"
const file = "log.txt"

func main() {

	//Timestamp in json
	tn := time.Now()
	ct := tn.Format("2006-01-02 15:04:05")
	replacer := strings.NewReplacer("2023-01-26T13:00:00+03:00", ct)
	jsonf := ""
	jsonf = replacer.Replace(jsonfe)

	f, err := os.ReadFile(jsonAdr)
	if err != nil {
		panic(err)
	}
	//Проверяем json
	value := gjson.Get(string(f), "Options.Create.Condition")
	value1 := gjson.Get(string(f), "Options.Change_Name.Condition")
	value2 := gjson.Get(string(f), "Options.Delete.Condition")
	value3 := gjson.Get(string(f), "Options.Get_Time.Condition")
	value4 := gjson.Get(string(f), "Options.Input.Condition")

	//Создаём файл
	if value.Num == 0 && value1.Num == 0 && value2.Num == 0 && value3.Num == 0 && value4.Num == 0 {
		f, err := os.OpenFile(file,
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		//Записываем произвольную строку
		timestamp := time.Now()
		fmt.Fprint(f, timestamp)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		//Меняем состояние json
		ChV, _ := sjson.Set(jsonf, "Options.Create.Condition", 1)

		z, err := os.OpenFile(jsonAdr, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		z.WriteString(ChV)

	}
	//Переименовываем файл
	if value.Num == 1 {

		err := os.Rename("log.txt", "avesome.txt")

		if err != nil {
			fmt.Println(err)
			return
		}
		//Вносим изменения в json
		a, _ := sjson.Set(jsonf, "Options.Change_Name.Condition", 1)
		z, err := os.OpenFile(jsonAdr, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		z.WriteString(a)
	}
	//Удаляем файл
	if value1.Num == 1 {
		err := os.Remove("avesome.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		//Вносим изменения в json
		b, _ := sjson.Set(jsonf, "Options.Delete.Condition", 1)
		z, err := os.OpenFile(jsonAdr, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		z.WriteString(b)

	}
	//Выводим время создания
	if value2.Num == 1 {
		g := gjson.Get(jsonAdr, "Options.Timestamp")
		log.Println(g)
		//Вносим изменения в json
		c, _ := sjson.Set(jsonf, "Options.Get_Time.Condition", 1)
		z, err := os.OpenFile(jsonAdr, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		z.WriteString(c)
	}
	g := gjson.Get(jsonAdr, "Options.Timestamp")
	//Проверяем время и переписываем json
	if value3.Num == 1 && g.Str < ct {
		//Вносим изменения в json
		d, _ := sjson.Set(jsonf, "Options.Get_Time.Condition", 0)
		z, err := os.OpenFile(jsonAdr, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			log.Fatal(err)
		}
		z.WriteString(d)
	}
}
