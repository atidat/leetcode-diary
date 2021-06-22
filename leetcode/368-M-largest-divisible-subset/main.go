package main

import (
	"fmt"
	"sort"
)

func main() {
	n1 := []int{1,2,3}
	fmt.Println(largestDivisibleSubset(n1)) // [1,2] or [1,3]
	fmt.Println()

	n2 := []int{1,2,4,8}
	fmt.Println(largestDivisibleSubset(n2)) // [1,2,4,8]
	fmt.Println()

	n3 := []int{1,3,6,7,14,28}
	fmt.Println(largestDivisibleSubset(n3)) // [1,7,14,28]

	n4 := []int{1,2}
	fmt.Println(largestDivisibleSubset(n4)) // [1,2] or [1,3]
	fmt.Println()

	n5 := []int{2,3,4,8}
	fmt.Println(largestDivisibleSubset(n5)) // [2,4,8]
}

type st struct {
	num int
	ind int
}

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make(map[int]st)

	head := 0
	if nums[0] == 1 { head = 1 }

	var m st
	for i := len(nums)-1; i >= head; i-- {
		j := i+1
		for ; j < len(nums); j++ {
			if nums[j] % nums[i] == 0 {
				dp[i] = st{dp[j].num+1, j}
				if m.num < dp[i].num {
					m = st{dp[i].num, j}
				}
				break
			}
		}
		if j == len(nums) {
			dp[i] = st{1, -1}
		}
	}

	fmt.Println(dp, m)
	res := make([]int, 0)
	if head == 1 {
		res = append(res, 1)
	}
	if m.ind == 0 && m.num == 0 {
		return append(res, nums[1])
	} else {
		m = st{m.num+1, m.ind-1}
	}


	i := m
	for ; i.ind != -1; {
		res = append(res, nums[i.ind])
		i = dp[i.ind]
	}
	return res
}
