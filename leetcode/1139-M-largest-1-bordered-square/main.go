package main

func main() {

}

func largest1BorderedSquare(grid [][]int) int {
	// row : 行； column： 列
	gr := len(grid)
	gl := len(grid[0])
	dp := make([][][2]int, gr)
	for i := range dp {
		dp[i] = make([][2]int, gl)
	}
	// dp[i][j][0]表示 grid[i][j][0]左边连续1的个数（含自身）
	// dp[i][j][1]表示 grid[i][j][0]上边连续1的个数（含自身）

	// 1 init dp
	for j := 0; j < gl; j++ {
		for i := 0; i < gr; i++ {
			if grid[i][j] == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = 0
				continue
			}
			if j-1 >= 0 {
				dp[i][j][0] = dp[i][j-1][0]
			}
			if i-1 >= 0 {
				dp[i][j][1] = dp[i-1][j][1]
			}

			if grid[i][j] == 1 {
				dp[i][j][0] += 1
				dp[i][j][1] += 1
			}
		}
	}

	//fmt.Println(dp)

	// 2 set dp
	curMax := 0
	res := 0
	for j := 0; j < gl; j++ {
		for i := 0; i < gr; i++ {
			curMax = mmin(dp[i][j][0], dp[i][j][1])
			for i-curMax+1 < gr && j-curMax+1 < gl {
				if dp[i-curMax+1][j][0] >= curMax && dp[i][j-curMax+1][1] >= curMax {
					break
				}
				curMax -= 1
			}
			res = mmax(res, curMax * curMax)
			//fmt.Println(res)
		}
	}
	return res
}

func mmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mmax(a, b int) int {
	if a < b {return b}
	return a
}