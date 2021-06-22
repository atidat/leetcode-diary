package main

import "fmt"

func main() {

}
/*
	0 < prices.length <= 50000.
	0 < prices[i] < 50000.
	0 <= fee < 50000.
*/
func maxProfit(prices []int, fee int) int {
/*	OLE
	dp := make([][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([]int, len(prices)+1)
	}

	for i := len(prices); i >= 1; i-- {
		for j := i+1; j <= len(prices); j++ {
			dp[i][j] = prices[j-1] - prices[i-1] - fee
			for k := i; k+1 <= j; k++ {
				dp[i][j] = mmax(dp[i][j], dp[i][k] + dp[k+1][j])
			}
		}
	}
	//fmt.Println(dp)
	return dp[1][len(prices)]*/
	//var res int
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	//dp[i][0] 表示i天不持股
	//dp[i][1] 表示i天持股
	dp[0][0], dp[0][1] = 0, 0-prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = mmax(dp[i-1][0], dp[i-1][1] + prices[i] - fee)
		dp[i][1] = mmax(dp[i-1][1], dp[i-1][0] - prices[i])
	}
	fmt.Println(dp)
	return mmax(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}
