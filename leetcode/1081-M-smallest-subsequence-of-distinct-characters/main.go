package main

import "fmt"

func main() {

}

func smallestSubsequence(s string) string {
	fmt.Println(s)
	visited := make(map[byte]int, 0)
	lastshow := make(map[byte]int, 0)
	for i := range s {
		lastshow[s[i]] = i
	}

	res := make([]byte, 0)
	for i := range s {
		_, ok1 := visited[s[i]]; if ok1 {
			continue
		}

		reslen := len(res)
		var cut = -1
		for j := reslen-1; j >= 0; j-- {
			// 后者替换前者唯一条件：后者值小；前者不是最后一次出现
			//fmt.Println(res)
			if i < lastshow[res[j]] && res[j] > s[i] {
				delete(visited, res[j])
				fmt.Println(visited[res[j]], lastshow[res[j]])
				cut = j
				//fmt.Println(cut)
			} else {
				break
			}
		}
		if cut >= 0 {
			res = res[:cut]
		}
		visited[s[i]] = i
		res = append(res,s[i])
		//fmt.Printf("%c\n", res)
	}
	//fmt.Printf("%c\n", res)
	return string(res)
}