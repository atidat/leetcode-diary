package main

import "fmt"

func main() {
	minCut("aab")
	minCut("a")
	minCut("ab")
	minCut("ababababababababababababcbabababababababababababa")
}

func minCut(s string) int {
	fmt.Println(s)
	dp := cache(s)
	cache := make([]int, len(s))
	for i := range cache {
		cache[i] = len(s)-1
	}
	cache[0] = 0
	return executor(dp, &cache)
}

func cache(src string) *[][]bool {
	sl := len(src)
	dp := make([][]bool, sl)
	for i := range dp {
		dp[i] = make([]bool, sl)
		dp[i][i] = true
	}
	for i := 0; i < sl-1; i++ {
		if src[i] == src[i+1] {
			dp[i][i+1] = true
		}
	}
	for l := 3; l <= sl; l++ {
		for s := 0; s < sl+1-l; s++ {
			e := s+l-1
			dp[s][e] = src[s] == src[e] && dp[s+1][e-1]
		}
	}
	return &dp
}

func executor(dp *[][]bool, ca *[]int) int {
	L := len(*ca)
	for i := 1; i < L; i++ {
		for k := 0; k+1 <= i; k++ {
			if (*dp)[k+1][i] {
				(*ca)[i] = mmin((*ca)[i], (*ca)[k]+1)
			}
		}

		m, n := 0, i
		for ; m <= n; {
			if !(*dp)[m][n] {break}
			m++
			n--
		}
		if m > n {
			(*ca)[i] = 0
		}


	}
	fmt.Println(*ca)
	return (*ca)[L-1]
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}
