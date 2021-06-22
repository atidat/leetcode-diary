package main

import "fmt"

func main() {
	fmt.Println(numTrees(6))
}

func numTrees(n int) int {
	if n == 1 || n ==2 {
		return n
	}


	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n; i++ {
		tmp := dynamic(i, dp)
		copy(dp, tmp)
		fmt.Println(dp)
	}
	sum := 0
	for i := range dp {
		sum += dp[i]
	}
	return sum
}

func dynamic(l int, dp []int) []int {
	tmp := make([]int, l+1)
	// [2,2,1]
	// [5,5,3,1]
	tmp[l] = 1
	for i := l-1; i > 0; i-- {
		for j := i-1; j < l; j++ {
			tmp[i] += dp[j]
		}
	}
	tmp[0] = tmp[1]
	return tmp
}