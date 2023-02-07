package actions

// CheckTime Сравнение времени с нынешним, для обновления цикла
func CheckTime(input JsonInput) {

	input.string = "Options.Get_Time.Condition"
	input.value = 0

	input.JsonWriter()
}
