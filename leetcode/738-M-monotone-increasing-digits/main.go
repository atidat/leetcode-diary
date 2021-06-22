package main

import (
	"fmt"
	"strconv"
)

func main() {

}

func monotoneIncreasingDigits(N int) int {
	tmpStr := strconv.Itoa(N)
	valb := []byte(tmpStr)
	return handle(valb)
}

//4031
func handle(valb []byte) int {
	vLen := len(valb)
	for i := vLen-1; i > 0; i-- {
		if valb[i] < valb[i-1] {
			valb[i] = '9'
			for j := i; j < vLen; j++ {
				valb[j] = '9'
			}
			if i == 1 && valb[i-1] == '1' {
				valb = valb[1:]
				break
			} else if valb[i-1] == '0' {
				valb[i-1] = '9'
			} else {
				valb[i-1] -= 1
			}
		}
	}
	res := calc(valb)
	fmt.Println(res)
	return res
}

func calc(v []byte) int {
	var res int
	for _, v := range v {
		res *= 10
		res += int(v - '0')
	}
	return res
}