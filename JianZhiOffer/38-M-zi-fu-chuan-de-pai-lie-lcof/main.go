package main

import (
	"sort"
)

func main() {
	permutation("abc")
	permutation("aba")
}

func permutation(s string) []string {
	opt := []byte(s)
	sort.Slice(opt, func(a, b int) bool {
		if opt[a] < opt[b] { return true }
		return false
	})

	rec := make([]bool, len(opt), len(opt))
	var res []string
	var mem []byte

	var recur func([]byte, int)
	recur = func (mem []byte, ind int) {
		if ind == len(opt) {
			res = append(res, string(mem))
			return
		}
		for i := 0; i < len(opt); i++ {
			if rec[i] {
				continue
			}
			if i-1 >= 0 && opt[i] == opt[i-1] && !rec[i-1] {
				continue
			}
			rec[i] = true
			mem = append(mem, opt[i])
			recur(mem, ind+1)
			mem = mem[:len(mem)-1]
			rec[i] = false
		}
	}
	recur(mem, 0)
	//fmt.Println(res)
	return res
}