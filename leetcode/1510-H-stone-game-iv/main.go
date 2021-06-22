package main

import "fmt"

func main() {

}

func winnerSquareGame(n int) bool {
	// dp[i] = true 表示 拿到这个数会赢
	dp := make([]bool, n+1)
	dp[0] = false
	dp[1] = true

	for i := 2; i <= n; i++ {
		dp[i] = false
		for j := 1; j*j <= n && i-j*j >= 0; j++ {
			if !dp[i-j*j] {
				dp[i] = true
				break
			}
		}
		//fmt.Println(i, dp)
	}
	fmt.Println(dp)
	return dp[n]
}
	/*if n == 1 { return true }
	if n == 2 { return false }

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = -1, -1, 1 // -1代表alice胜利  1代表bob胜利
	for i := 3; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = opt(dp[i], dp[j*j] * dp[i-j*j])
		}
	}
	return dp[n] == -1
}

func opt(a, b int) int {
	if a == -1 || b == -1 {
		return -1
	}
	return 1
}*/