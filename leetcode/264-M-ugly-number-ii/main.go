package main

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	ind2, ind3, ind5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = getMin(2*dp[ind2], getMin(3*dp[ind3], 5*dp[ind5]))
		if dp[i] == 2*dp[ind2] {
			ind2++
		}
		if dp[i] == 3*dp[ind3] {
			ind3++
		}
		if dp[i] == 5*dp[ind5] {
			ind5++
		}
	}

	return dp[n]
}

func main() {
	//_ = nthUglyNumber(10)
	_ = nthUglyNumber(1690)
}
