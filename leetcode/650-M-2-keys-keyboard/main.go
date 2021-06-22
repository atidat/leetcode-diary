package main

func main() {

}

var prime = []int{2,3,5,7,11,13,17,19,23,29,31}
/*
n 的取值范围是 [1, 1000] 。
恰好打印出 n个字符
*/
func minSteps(n int) int {
	if n == 1 { return 0 }

	dp := make([]int, n+1)
	dp[1], dp[2] = 0, 2
	for t := 3; t <= n; t++ {
		dp[t] = t
		for i := 0; i < len(prime); i++ {
			if t%prime[i] == 0 {
				dp[t] = mmin(dp[t], dp[t/prime[i]]+prime[i])
			}
		}
	}
	return dp[n]
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}