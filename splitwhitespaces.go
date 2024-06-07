package piscine

func SplitWhiteSpaces(s string) []string {
	var word string
	var result []string
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] != ' ' && runes[i] != '\t' && runes[i] != '\n' {
			word += string(runes[i])
		} else {
			result = append(result, word)
			word = ""
		}
		if i == len(runes)-1 && runes[i] != ' ' && runes[i] != '\t' && runes[i] != '\n' {
			result = append(result, word)
		}
	}
	return result
}
