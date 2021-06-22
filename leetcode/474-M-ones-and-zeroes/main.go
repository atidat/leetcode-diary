package main

import (
	"fmt"
)

func main() {

}

/*
	1 <= strs.length <= 600
	1 <= strs[i].length <= 100
	strs[i] 仅由 '0' 和 '1' 组成
	1 <= m, n <= 100
*/
type st struct {
	z int
	o int
}

// m:0   n:1
func findMaxForm(strs []string, m int, n int) int {
	mm := mapping(strs)
	fmt.Println(strs)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for h := range strs {
		t := mm[strs[h]]
		for i := m; i >= t.z; i-- {
			for j := n; j >= t.o; j-- {
				dp[i][j] = mmax(dp[i-t.z][j-t.o]+1, dp[i][j])
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

func mapping(strs []string) map[string]st {
	m := make(map[string]st, 0)
	for i := range strs {
		m[strs[i]] = calc(strs[i])
	}
	return m
}

func calc(s string) st {
	zero, one := 0, 0
	for i := range s {
		if s[i] == '1' {
			one++
		} else {
			zero++
		}
	}
	return st{zero, one}
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}