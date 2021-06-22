package main

import "fmt"

func main() {

}

/*
如：d = [1,17,5,10,13,15,10,5,16,8]
求d[0] ~ d[n-1]的摇摆序列，其实是求
	d[0] ~ d[n-2] 的摇摆序列的最大值，是否加1，取决于d[n-2]和其前后的差值乘积是否 < 0

*/
func wiggleMaxLength(nums []int) int {
	numLen := len(nums)
	if numLen < 2 { return numLen }
	res := executor(nums)
	fmt.Println(res)
	return res
}


func executor(nums []int) int {
	/*该dp变化量是趋势，即（当前值-上一个值）是 + or - */
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, len(nums))
		dp[i][0] = 1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i] -nums[i-1] > 0 {
			dp[0][i] = mmax(dp[1][i-1]+1, dp[0][i-1])
		} else if nums[i] - nums[i-1] < 0 {
			dp[1][i] = mmax(dp[0][i-1]+1, dp[1][i-1])
		} else {
			dp[0][i] = dp[0][i-1]
			dp[1][i] = dp[1][i-1]
		}
	}
	//fmt.Println(dp)
	return mmax(dp[0][len(nums)-1], dp[1][len(nums)-1])
}

func mmax(a, b int) int {
	if a < b {
		return b
	}
	return a
}