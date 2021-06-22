package main

var alpha [27]string = [27]string{"", "A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

func recur(n int, res string) string {
	de := 26
	if n < 27 {
		return res + alpha[n]
	}

	if n%de == 0 {
		return recur(n/de-1, res) + alpha[26]
	}
	return recur(n/de, res) + alpha[n%de]
}

func convertToTitle(n int) string {
	return recur(n, "")
}

func main() {
	_ = convertToTitle(703)
}
