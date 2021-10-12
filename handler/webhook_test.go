package handler

func getRandomFields() []string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	var fields []string
	for i := 0; i < 3; i++ {
		fields = append(fields, string(letters[i]))
	}
	return fields
}
