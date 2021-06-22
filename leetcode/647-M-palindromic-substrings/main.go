package main

import "fmt"

func main() {

}

func countSubstrings(s string) int {
	sl := len(s)
	dp := make([][]int, sl)
	for i := range dp {
		dp[i] = make([]int, sl)
		dp[i][i] = 1
	}

	ca := cache(s)
	for i := sl-1; i >= 0; i-- {
		for j := i+1; j < sl; j++ {
			dp[i][j] = 1 + dp[i+1][j]

			for k := i+1; k < sl; k++ {
				if (*ca)[i][k] {
					dp[i][j] += 1
				}
			}
			//fmt.Println(dp)
		}
	}
	fmt.Println(dp[0][sl-1])
	return dp[0][sl-1]
}

func cache(s string) *[][]bool {
	sl := len(s)
	ca := make([][]bool, sl)
	for i := range ca {
		ca[i] = make([]bool, sl)
	}

	for i := sl-1; i >= 0; i-- {
		ca[i][i] = true
		for j := i+1; j < sl; j++ {
			h, t := i, j
			for ; h < t; {
				if s[h] != s[t] {
					ca[i][j] = false
					break
				}
				h++; t--
			}
			if h >= t {
				ca[i][j] = true
			}
		}
	}
	return &ca
}
