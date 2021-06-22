package main

import "fmt"

func main() {
	fmt.Println(findSubstringInWraproundString("abcdhijcdefgh"))
}

func findSubstringInWraproundString(p string) int {
	p += "#"
	/* 计算子串，可以理解为以母串中各字符结尾的串 */
	repo := make(map[byte]int)

	var conti_len = 1
	for i := 1; i < len(p); i++ {

		v, ok := repo[p[i-1]]
		if !ok || conti_len > v {
			repo[p[i-1]] = conti_len
		}

		if p[i-1] + 1 == p[i] || p[i] == 'a' && p[i-1] == 'z' {
			conti_len += 1
		} else {
			conti_len = 1
		}
	}

	var res int
	for _, v := range repo {
		//fmt.Printf("%c, %d\n", k, v)
		res += v
	}
	return res
}