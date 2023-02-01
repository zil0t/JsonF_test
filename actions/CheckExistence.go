package actions

import "os"

// Existence Проверка существует ли файл
func Existence(f string, err error) {
	_, err = os.Stat(f)
	if err != nil {
		panic("Файла нет")
	}
}
