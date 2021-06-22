package main

import (
	"sort"
)

func main() {

}

func combinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	if len(nums) == 0 || target < nums[0] { return 0 }
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = -1
	}
	return exec(nums, target, &dp)

}

func exec(nums []int, target int, dp *[]int) int {
	if (*dp)[target] != -1 {
		return (*dp)[target]
	}
	var res int
	for i := range nums {
		if target < nums[i] { break }
		if target == nums[i] {
			res += 1
			break
		}
		if target > nums[i] {
			res += exec(nums, target-nums[i], dp)
			continue
		}
	}
	(*dp)[target] = res
	return res
}
