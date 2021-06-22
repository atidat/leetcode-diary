package main

import "fmt"

func phoneDict() map[byte][]byte {
	phone := make(map[byte][]byte, 0)
	phone['2'] = []byte{'a', 'b', 'c'}
	phone['3'] = []byte{'d', 'e', 'f'}
	phone['4'] = []byte{'g', 'h', 'i'}
	phone['5'] = []byte{'j', 'k', 'l'}
	phone['6'] = []byte{'m', 'n', 'o'}
	phone['7'] = []byte{'p', 'q', 'r', 's'}
	phone['8'] = []byte{'t', 'u', 'v'}
	phone['9'] = []byte{'w', 'x', 'y', 'z'}
	return phone
}
func iterator(digits string, i int, p map[byte][]byte, repo *[]byte, dict *[]string) {
	if i == len(digits) {
		*dict = append(*dict, string(*repo))
		return
	}
	for j := 0; j < len(p[digits[i]]); j++ {
		*repo = append(*repo, p[digits[i]][j])
		iterator(digits, i+1, p, repo, dict)
		*repo = (*repo)[0:len(*repo)-1]
	}
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	p := phoneDict()
	dict := make([]string, 0)
	tmp := make([]byte, 0)
	iterator(digits, 0, p, &tmp, &dict)
	return dict
}

func main() {
	fmt.Println(letterCombinations("4576"))
}
