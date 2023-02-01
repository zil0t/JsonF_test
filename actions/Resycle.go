package actions

// CheckTime Сравнение времени с нынешним, для обновления цикла
func CheckTime(json string) {
	//Вносим изменения в json
	input := JsonInput{
		json,
		"Options.Get_Time.Condition",
		0,
	}

	input.JsonWriter()
}
