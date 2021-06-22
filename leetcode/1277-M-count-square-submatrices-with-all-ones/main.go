package main

func main() {

}

func countSquares(matrix [][]int) int {
	//fmt.Println(matrix)
	ml, mml := len(matrix), len(matrix[0])

	var sum int

	dp := make([][]int, ml)
	for i := range dp {
		dp[i] = make([]int, mml)
		dp[i][0] = matrix[i][0]
		sum += dp[i][0]
	}
	for j := 1; j < mml; j++ {
		dp[0][j] = matrix[0][j]
		sum += dp[0][j]
	}


	for i := 1; i < ml; i++ {
		for j := 1; j < mml; j++ {
			if matrix[i][j] == 0 { dp[i][j] = 0; continue }
			b, s := maxmin(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			if b == s { dp[i][j] = b + 1 } else { dp[i][j] = s+1 }
			sum += dp[i][j]
		}
	}
	//fmt.Println(dp, sum)  // TODO debug
	return sum
}

func maxmin(a, b, c int) (int, int) {
	if a > b {
		if b > c { return a, c }
		if a > c { return a, b } else { return c, b }
	} else { // a < b
		if a > c { return b, c}
		if b > c { return b, a } else { return c, a }
	}
}