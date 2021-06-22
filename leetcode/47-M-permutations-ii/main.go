package main

import (
	"fmt"
	"sort"
)

func main() {

}

/*
*  1 <= nums.length <= 8
 -10 <= nums[i] <= 10
*/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	val := make([]int, 0)
	only := make(map[int]bool, 0)
	rec := make(map[int]bool)
	dfs(nums, rec, &res, val, only)
	fmt.Println(res)
	return res
}

func dfs(nums []int, rec map[int]bool, res *[][]int, val []int, only map[int]bool) {
	nl := len(nums)
	if len(val) == nl {
		h := handler(val)
		if _, ok := only[h]; ok {
			return
		}

		only[h] = true
		tmp := make([]int, nl)
		copy(tmp, val)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < nl; i++ {
		_, ok := rec[i]
		if ok {
			continue
		}

		val = append(val, nums[i])
		rec[i] = true
		dfs(nums, rec, res, val, only)
		delete(rec, i)
		val = val[:len(val)-1]
	}
}

func handler(v []int) int {
	var res int
	for _, v := range v {
		res *= 10
		res += v
	}
	return res
}