package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
	dp [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{matrix: matrix, dp: matrix}
	}

	lenM := len(matrix)+1
	lenM0 := len(matrix[0])+1
	dp := make([][]int, lenM)
	dp[0] = make([]int, lenM0)

	for i := 1; i < lenM; i++ {
		dp[i] = make([]int, lenM0)
		for j := 1; j < lenM0; j++ {
			if j == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
			}
		}
	}
	fmt.Println(dp)
	return NumMatrix{matrix: matrix, dp: dp}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.dp[row2+1][col2+1] - this.dp[row1+1][col1+1] + this.matrix[row1][col1] -
		(this.dp[row1][col2+1] - this.dp[row1][col1+1]) -
		(this.dp[row2+1][col1] - this.dp[row1+1][col1])
}

func main() {
}
