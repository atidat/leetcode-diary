package main

import "fmt"

//import "fmt"

func main() {

}
/*
func stoneGame(piles []int) bool {
	// dp[i][j] 表示 [i, j] 当前先手拿到的差值
	dp := cache(len(piles))
	return executor(&piles, dp, 0, len(piles)-1) > 0
}

func cache(pl int) *[][]int {
	dp := make([][]int, pl)
	for i := range dp {
		dp[i] = make([]int, pl)
	}
	return &dp
}

func executor(p *[]int, dp *[][]int, h, t int) int {
	if (*dp)[h][t] != 0 {
		return (*dp)[h][t]
	}

	var res int
	res = mmax(res, (*p)[h] - (*dp)[h+1][t])
	res = mmax(res, (*p)[t] - (*dp)[h][t-1])
	(*dp)[h][t] = res
	return res
}


func mmax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}*/

func stoneGame(piles []int) bool {
	mmax := func(a, b int) int {
		if a < b { return b }
		return a
	}
	/*
	每个人都采取最优策略
	每次只能从头或尾拿
	*/
	pl := len(piles)
	dp := make([][]int, pl)
	for i := 0; i < pl; i++ {
		dp[i] = make([]int, pl)
		dp[i][i] = piles[i]
	}
	for i := pl; i >= 0; i-- {
		for j := i+1; j < pl; j++ {
			dp[i][j] = mmax(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])
		}
	}
	fmt.Println(dp)
	return dp[0][pl-1] > 0
}
