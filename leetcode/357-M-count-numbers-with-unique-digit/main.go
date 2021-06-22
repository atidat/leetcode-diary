package main
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	total := 1
	inval := -9
	for n != 0 {
		total *= 10
		inval += 9
		n--
	}
	return total - inval
}

func main() {
}
