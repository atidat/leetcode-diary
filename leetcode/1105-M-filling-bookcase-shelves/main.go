package main

import (
	"fmt"
	"math"
)

func main() {

}
func minHeightShelves(books [][]int, shelf_width int) int {
	dp := make([]int, len(books)+1)
	for i := 1; i <= len(books); i++ {
		dp[i] = math.MaxInt32
	}
	// books[i][j]  i:厚度  j:高度
	for k := 1; k <= len(books); k++ {
		aw := books[k-1][0]
		maxh := books[k-1][1]
		for t := k-1; t >= 0 && aw <= shelf_width; {
			dp[k] = mmin(dp[k], maxh + dp[t])
			if t == 0 {
				break
			}
			t--
			aw += books[t][0]
			maxh = mmax(maxh, books[t][1])
		}
	}
	fmt.Println(dp)
	return dp[len(books)]

}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}