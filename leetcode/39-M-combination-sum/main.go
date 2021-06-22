package main

import (
	"fmt"
	"sort"
)

func main() {
	combinationSum([]int{2,3,6,7}, 7)
	fmt.Printf("expect: [[2,2,3], [7]]\n")
	combinationSum([]int{2,3,5}, 8)
	fmt.Printf("expect: [[2,2,2,2], [2,3,3], [3,5]]\n")
	combinationSum([]int{2}, 1)
	fmt.Printf("expect: []\n")
	combinationSum([]int{1}, 1)
	fmt.Printf("expect: [[1]]\n")
	combinationSum([]int{1}, 2)
	fmt.Printf("expect: [[1,1]]\n")
	combinationSum([]int{5,8}, 3)
	fmt.Printf("expect: []\n")
	combinationSum([]int{2,3,5}, 8)
	fmt.Printf("expect: [[2,2,2,2],[2,3,3],[3,5]]\n")
}

func combinationSum(c []int, t int) [][]int {
	sort.Ints(c)
	val := make([]int, 0)
	res := make([][]int, 0)
	dfs(c, &res, t, val, 0)
	fmt.Println(res)
	return res
}

func dfs(c []int, res *[][]int, t int, val []int, ind int) {
	if t == 0 {
		t := make([]int, len(val))
		copy(t, val)
		*res = append(*res, t)
		return
	}

	cl := len(c)
	for i := ind; i < cl; i++ {
		if t - c[i] < 0 {
			break
		}
		val = append(val, c[i])
		dfs(c, res, t-c[i], val, i)
		val = val[:len(val)-1]
	}
}
