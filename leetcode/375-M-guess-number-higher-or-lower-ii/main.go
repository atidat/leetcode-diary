package main

import (
	"fmt"
	"math"
)

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][i] = 0
		for j := i+1; j <= n; j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	for i := n-1; i >= 1; i-- {
		for j := i+1; j <= n; j++ {
			for k := i+2; k+1 <= j; k++ {
				dp[i][j] = mmin(dp[i][j], dp[i][k-1] + k + dp[k+1][j])
			}
			dp[i][j] = mmin(dp[i][j], i+dp[i+1][j])
			dp[i][j] = mmin(dp[i][j], dp[i][j-1]+j)
		}
	}
	fmt.Println(dp)
	return dp[1][n]
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}

func main() {
	getMoneyAmount(5)
}
