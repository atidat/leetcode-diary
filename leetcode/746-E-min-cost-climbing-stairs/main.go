package main

import "fmt"

func main() {

}

func minCostClimbingStairs(cost []int) int {
	cl := len(cost)
	dp := make([]int, cl)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < cl; i++ {
		dp[i] = mmin(dp[i-1], dp[i-2]) + cost[i]
	}
	fmt.Println(cost)
	fmt.Println(dp)
	return mmin(dp[cl-1], dp[cl-2])
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}