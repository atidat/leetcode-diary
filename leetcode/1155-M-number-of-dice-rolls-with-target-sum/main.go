package main

func main() {

}

/*
这里有 d 个一样的骰子，每个骰子上都有 f 个面，分别标号为 1, 2, ..., f。

我们约定：掷骰子的得到总点数为各骰子面朝上的数字的总和。

如果需要掷出的总点数为 target，请你计算出有多少种不同的组合情况（
所有的组合情况总共有 f^d 种），模 10^9 + 7 后返回。


  1 <= d, f <= 30
  1 <= target <= 1000
*/
func numRollsToTarget(d int, f int, target int) int {
	// d 表示骰子数 f 表示骰子面数  target 表示目标和
	dp := make([][]int, d+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = 0
		}
	}
	dp[0][0] = 1

	for i := 1; i <= d; i++ {
		for j := i; j <= target; j++ {
			for k := 1; k <= f ; k++ {
				if j >= k {
					dp[i][j] += dp[i-1][j-k] % (1e9 + 7)
				}
			}
		}
	}
	//fmt.Println(dp)
	// dp[i][tar] = dp[i-1][tar-1] + ... dp[i-1][tar-f]
	return dp[d][target] % (1e9 + 7)
}
