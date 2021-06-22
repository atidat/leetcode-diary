package main

import (
	"math"
)

func main() {

}

/*
  1 <= cardPoints.length <= 10^5
  1 <= cardPoints[i] <= 10^4
  1 <= k <= cardPoints.length
*/
func maxScore(cardPoints []int, k int) int {

	cpl := len(cardPoints)
	dp := make([]int, cpl+1)
	for j := 1; j <= cpl; j++ {
		dp[j] = dp[j-1] + cardPoints[j-1]
	}

	res := math.MinInt32
	fstRun, sndRun := k, cpl
	for fstRun >= 0 {
		res = mmax(res, dp[fstRun] + dp[cpl] - dp[sndRun])
		fstRun -= 1
		sndRun -= 1
	}
	return res
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}