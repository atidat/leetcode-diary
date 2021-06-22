package main

import "fmt"

func main() {

}

var trail = [][]int{
	{-2, 1}, {-1, 2},
	{1, 2},	{2, 1},
	{2, -1}, {1,-2},
	{-1, -2}, {-2, -1},
}

/*
solution 1: TLE
func knightProbability(N int, K int, r int, c int) float64 {
	succ, total := exec(N, K, r, c)
	fmt.Println(succ, total)
	return float64(succ)/float64(total)
}

func exec(scope int, times int, r, c int) (int, int) {
	var succ, total int

	if r >= scope || r < 0 || c >= scope || c < 0 {
		return succ, pow(times)
	}
	if times == 0 {
		return 1, 1
	}

	for i := 0; i < 8; i++ {
		var succ_tmp, total_tmp int
		succ_tmp, total_tmp = exec(scope, times-1, r+trail[i][0], c+trail[i][1])

		succ += succ_tmp
		total += total_tmp
	}
	return succ, total
}

func pow(b int) int {
	res := 1
	for i := 1; i <= b; i++ {
		res <<= 3
	}
	return res
}*/
func knightProbability(N int, K int, r int, c int) float64 {
	dp := make([][][]float64, K+1)
	for t := range dp {
		dp[t] = make([][]float64, N)
		for i := range dp[t] {
			dp[t][i] = make([]float64, N)
		}
	}
	for i := range dp[0] {
		for j := range dp[0][i] {
			dp[0][i][j] = 1.0
		}
	}

	for t := 1; t <= K; t++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				for h := 0; h < 8; h++ {
					iRun, jRun := i+trail[h][0], j+trail[h][1]
					if iRun >= N || iRun < 0 || jRun >= N || jRun < 0 {
						//
					} else {
						dp[t][i][j] += dp[t-1][iRun][jRun]/8.0
					}
				}
			}
		}
	}
	fmt.Println(dp[K][r][c])
	return dp[K][r][c]
}

func pow(times int) int {
	res := 1
	for i := 1; i <= times; i++ {
		res <<= 3
	}
	return res
}