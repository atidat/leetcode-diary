package main

import (
	"fmt"
	"sort"
)

func main() {

}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	r := make([]int, 0)
	res := make([][]int, 0)
	handler(candidates, target, 0, r, &res)
	fmt.Println(res)
	return res
}

func handler(c []int, tar int, ind int, r []int, res *[][]int) {
	if tar == 0 {
		*res = append(*res, r)
		return
	}

	cl := len(c)
	for i := ind; i < cl; i++ {
		if tar - c[i] < 0 {
			break
		}
		if i > ind && c[i] == c[i-1] {
			continue
		}
		r := append(r, c[i])
		handler(c, tar-c[i], i+1, r, res)
		r = r[:len(r)-1]
	}
}
