// Quest9
package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, v := range tab {
		if f(v) {
			count++
		}
	}
	return count
}
