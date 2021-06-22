package main

import (
	"fmt"
	"strings"
)

func main() {

}

func isPowerOfFour(n int) bool {
	nbstr := fmt.Sprintf("%b", n)
	fmt.Println(nbstr)
	t := strings.Count(nbstr, "1")
	return t == 1 && nbstr[0] == '1' && len(nbstr) % 2 == 1
}