package piscine

func Join(strs []string, sep string) string {
	var str string
	for i := 0; i < len(strs); i++ {
		str += string(strs[i]) + sep + " "
	}
	return str
}
