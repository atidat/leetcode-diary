package main

import "fmt"

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func rob(nums []int) int {
	fmt.Println(nums, " ======")
	lenNums := len(nums)
	if lenNums == 0 {
		return 0
	}

	dp := make([]int, lenNums+1)
	dp[0] = 0
	dp[1] = nums[0]

	for k := 1; k < lenNums; k++ {
		dp[k+1] = getMax(dp[k], dp[k-1]+nums[k])

		//fmt.Println(hist, dp)
	}

	return dp[lenNums]
}
