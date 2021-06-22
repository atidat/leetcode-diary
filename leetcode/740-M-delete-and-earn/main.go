package main

import (
	"fmt"
	"sort"
)

func main() {

}

/*
  1 <= nums.length <= 2 * 104
  1 <= nums[i] <= 104
*/
func deleteAndEarn(in []int) int {
	sort.Ints(in)
	nums := make([]int, 0)
	m := make(map[int]int, 0)
	for i := range in {
		if _, ok := m[in[i]]; !ok {
			m[in[i]] = 1
			nums = append(nums, in[i])
		} else {
			m[in[i]] += 1
		}
	}

	nl := len(nums)
	dp := make([][]int, nl)
	for i := range dp {
		dp[i] = make([]int, nl)
		dp[i][i] = nums[i] * m[nums[i]]
	}

	for i := nl-1; i >= 0; i-- {
		for j := i+1; j < nl; j++ {
			for k := i; k+1 <= j; k++ {
				tmp := nums[k] * m[nums[k]]
				if k-1 < 0 {
					tmp += 0
				} else if nums[k-1] + 1 != nums[k] {
					tmp += dp[i][k-1]
				} else if i <= k-2 {
					tmp += dp[i][k-2]
				}

				if k+1 > nl-1 {
					tmp += 0
				} else	if nums[k] + 1 != nums[k+1] {
					tmp += dp[k+1][j]
				} else if k+1 <= j {
					tmp += dp[k+1][j]
				}
				dp[i][j] = mmax(dp[i][j], tmp)
			}
		}
	}
	fmt.Println(dp)
	return dp[0][nl-1]
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}