package piscine

func IterativePower(nb int, power int) int {
	result := 1
	pw := power
	num := nb

	for power != 0 {
		result *= num
		pw -= 1
	}
	if result < 0 {
		return 0
	}
	return result
}
