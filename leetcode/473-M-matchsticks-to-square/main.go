package main

import (
	"fmt"
	"sort"
)

func main() {
	if makesquare([]int{3, 3, 3, 3, 4}) {
		fmt.Println("bingo")
	} else {
		fmt.Println("nooooo")
	}
}

func makesquare(nums []int) bool {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	total := 0
	for _, v := range nums {
		total += v
	}

	if total%4 != 0 {
		fmt.Println("1")
		return false
	}
	ave := total / 4

	sums := make([]int, 4)
	for i := 0; i < 4; i++ {
		sums[i] = ave
	}

	return dfs(&sums, nums, 0)
}

func dfs(sums *[]int, nums []int, i int) bool {
	nLen := len(nums)
	if i == nLen {
		return true
	}

	for j := 0; j < 4; j++ {
		if (*sums)[j] >= nums[i] {
			(*sums)[j] -= nums[i]
			if dfs(sums, nums, i+1) {
				return true
			}
			(*sums)[j] += nums[i]
		}
		//fmt.Println(*sums)
	}

	return false
}
