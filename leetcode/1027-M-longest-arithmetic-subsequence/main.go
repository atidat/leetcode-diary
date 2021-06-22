package main

func main() {

}
/*
  2 <= A.length <= 2000
  0 <= A[i] <= 10000
*/

func longestArithSeqLength(arr []int) int {
	// 求最长等差子序列
	al := len(arr)
	res := 2
	dp := make([][]int, al)
	for i := range dp {
		dp[i] = make([]int, al)
		for j := range dp[i] {
			dp[i][j] = 2
		}
	}

	m := make(map[int][]int)
	for k, v := range arr {
		if val, ok := m[v]; ok {
			m[v] = append(val, k)
		} else {
			m[v] = []int{k}
		}
	}

	for i := 0; i < al; i++ {
		for j := i+1; j < al; j++ {
			if v, ok := m[arr[i] + arr[i] - arr[j]]; ok {
				for k := range v {
					if v[k] < i {
						dp[i][j] = mmax(dp[i][j], dp[v[k]][i] + 1)
					}

				}
			}
			res = mmax(res, dp[i][j])
		}
	}
	return res
}


func mmax(a, b int) int {
	if a < b { return b }
	return a
}