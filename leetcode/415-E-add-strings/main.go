package main

import "fmt"

func main() {
	//addStrings("1123456789", "987654321")
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("11", "99"))
	fmt.Println(addStrings("1", "9"))
	fmt.Println(addStrings("9", "99"))
	fmt.Println(addStrings("5412","4841"))
	fmt.Println(addStrings("0", "0"))
}

func addStrings(num1 string, num2 string) string {
	long, short := num1, num2
	if len(num1) < len(num2) {
		long, short = short, long
	}
	long = "0" + long

	var b []byte
	for i := 0; i < len(long) - len(short); i++ {
		b = append(b, '0')
	}
	short = string(b) + short
	res := make([]byte, len(long), len(long))
	carry := 0
	for i := len(long)-1; i >= 0; i-- {
		tmp := int(long[i] - '0') + int(short[i] - '0') + carry
		if tmp > 9 {
			res[i] = byte(tmp - 10 + '0')
			carry = 1
		} else {
			res[i] = byte(tmp + '0')
			carry = 0
		}
	}
	if res[0] == '0' {
		return string(res[1:])
	}
	return string(res)
}