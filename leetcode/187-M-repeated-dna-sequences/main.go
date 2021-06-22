package main

import (
	"fmt"
)

func findRepeatedDnaSequences(s string) []string {
	lens := len(s)
	if lens <= 10 {
		return []string{}

	codemtd := 0
	for i := 0; i < 9; i++ {
		codemtd = codemtd<<3 | (int(s[i]) & 7)
	}
	repo := make(map[int]int)
	res := make([]string, 0)
	for i := 9; i < lens; i++ {
		codemtd = (codemtd&0x7ffffff)<<3 | (int(s[i]) & 7)
		val, ok := repo[codemtd]
		if !ok {
			repo[codemtd] = 0
		}
		if ok && val == 0 {
			res = append(res, s[i-9:i+1])
			repo[codemtd]++
		}
	}
	return res
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
