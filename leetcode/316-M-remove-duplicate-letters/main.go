package main

func main() {

}

// 去除重复字母，且要求字典序最小
func removeDuplicateLetters(s string) string {
	dict := new([256]int)
	for k := range s {
		dict[s[k]]++
	}
	// fmt.Println(dict)

	vis := new([256]bool)
	res := make([]byte, 1)
	res[0] = '0'

	for k := range s {
		dict[s[k]]--
		if vis[s[k]] {
			continue
		}
		for s[k] < res[len(res)-1] && dict[res[len(res)-1]] > 0 {
			vis[res[len(res)-1]] = false
			res = res[:len(res)-1]
		}
		res = append(res, s[k])
		vis[s[k]] = true
	//	fmt.Printf("%c\n", res)
	}

	res = res[1:]
	//fmt.Println(string(res))
	return string(res)
}
