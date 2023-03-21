package iteration

// Method to repeat a given string 4 times
func Repeat(text string) string {
	repeatedText := ""
	for index := 0; index < 4; index++ {
		repeatedText += text
	}
	return repeatedText
}
