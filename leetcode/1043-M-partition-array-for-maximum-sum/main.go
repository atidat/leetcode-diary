package main

func main() {

}

func maxSumAfterPartitioning(arr []int, sp int) int {
	al := len(arr)
	dp := make([][][]int, sp+1)
	for i := range dp {
		dp[i] = make([][]int, al+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, al+1)
		}
	}

	for k := 1; k <= sp; k++ {
		for i := 1; i <= al; i++ {
			for j := 1; j <= al; j++ {
				for t := 1; t < k; t++ {
					for m := i; m <= j; m++ {
						if m-i+1 >= t && j-m >= t-k {
							dp[k][i][j] = dp[t][i][m] + dp[t-k][m+1][j]
						}							
					}

				}
			}
		}
	}
	return dp[sp][1][al]
}