package main

import (
	"fmt"
	"math"
)

func main() {

}

/*
	3 <= A.length <= 50
	1 <= A[i] <= 100
*/
func minScoreTriangulation(arr []int) int {
	dp := make([][]int, len(arr))
	for i := range dp {
		dp[i] = make([]int, len(arr))
	}

	for t := 2; t < len(arr); t++ {
		for i := 0; i+t < len(arr); i++ {
			j := i+t

			if t == 2 {
				dp[i][j] = arr[i] * arr[i+1] * arr[j]
			} else {
				dp[i][j] = math.MaxInt32
				for k := i+1; k < j; k++ {
					dp[i][j] = mmin(dp[i][j], arr[i] * arr[k] * arr[j] + dp[i][k] + dp[k][j])
				}
			}
		}
	}
	fmt.Println(dp[0][len(arr)-1])
	return dp[0][len(arr)-1]
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}
