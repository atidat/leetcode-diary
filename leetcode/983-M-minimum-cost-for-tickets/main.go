package main

import (
	"fmt"
	"math"
)

func main() {

}

func mincostTickets(days []int, costs []int) int {
	final := days[len(days)-1]
	dp := make([]int, final+1)
	for i := 1; i <= final; i++ {
		dp[i] = math.MaxInt32
	}
	scope := make(map[int]int, 0)
	scope[costs[0]] = 1
	scope[costs[1]] = 7
	scope[costs[2]] = 30
	for k, j := 1, 0; k <= final; k++ {
		if k != days[j] {
			dp[k] = dp[k-1]
			continue
		}


		for i := 0; i < 3; i++ {
			if k-scope[costs[i]] >= 0 {
				dp[k] = mmin(dp[k], dp[k-scope[costs[i]]]+costs[i])
			} else {
				dp[k] = mmin(dp[k], dp[0]+costs[i])
			}
		}
		//fmt.Println(dp[1:])
		j++
	}
	fmt.Println(dp[1:])
	return dp[final]
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}