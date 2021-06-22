package main

import "fmt"

func getMax(a, b int) int {
	if a <  b {
		return b
	}
	return a
}

func integerBreak(n int) int {
	if n == 2 || n == 3 {
		return n - 1
	}

	dp := make([]int, n+1)
	dp[1], dp[2], dp[3] = 1, 2, 3
	for k := 4; k < n + 1; k++ {
		for i := 1; i <= k/2; i++ {
			dp[k] = getMax(dp[k], dp[i] * dp[k-i])
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(integerBreak(10))
}
