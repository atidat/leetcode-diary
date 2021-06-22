package main

import "math"

func main() {

}
func orderOfLargestPlusSign(N int, mines [][]int) int {
	// construct data source
	points := make([][]int, N)
	for i := range points {
		points[i] = make([]int, N)
		for j := 0; j < N; j++ {
			points[i][j] = 1
		}
	}
	if len(mines) != 0 && len(mines[0]) != 0 {
		for i := range mines {
			points[mines[i][0]][mines[i][1]] = 0
		}
	}

	//fmt.Println(points)  // TODO
	// dp[i][j][0] 表示 points[i][j]上边有 xx 个连续1
	// dp[i][j][1] 表示 points[i][j]右边有 xx 个连续1
	// dp[i][j][2] 表示 points[i][j]下边有 xx 个连续1
	// dp[i][j][3] 表示 points[i][j]左边有 xx 个连续1
	dp := make([][][]int, N+2)
	for i := range dp {
		dp[i] = make([][]int, N+2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 4)
		}
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if points[i-1][j-1] != 0 {
				dp[i][j][0] = dp[i-1][j][0] + 1
				dp[i][j][3] = dp[i][j-1][3] + 1
			}

			if points[N-i][N-j] != 0 {
				dp[N+1-i][N+1-j][1] = dp[N+1-i][N+1-j+1][1] + 1
				dp[N+1-i][N+1-j][2] = dp[N+1-i+1][N+1-j][2] + 1
			}
		}
	}

	res := math.MinInt32
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			tmp := math.MaxInt32
			for k := 0; k < 4; k++ {
				tmp = mmin(tmp, dp[i][j][k])
			}
			res = mmax(res, tmp)
		}
	}
	return res
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}