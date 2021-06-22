package main

import "fmt"

func main() {

}

func largestSumOfAverages(A []int, K int) float64 {
	accu := calc(&A)
	dp := make([][]float64, K+1)
	for i := range dp {
		dp[i] = make([]float64, len(A))
	}
	res := exec(&dp, accu, K, len(A))
	fmt.Println(res)
	return res
}

func exec(dp *[][]float64, accu *[]int, K, ll int) float64 {
	if (*dp)[K][ll-1] > 0 {
		return (*dp)[K][ll-1]
	}
	if K <= 1 {
		return float64((*accu)[ll-1])/float64(ll)
	}

	// dp[K][ll] 表示 前ll个元素，划分成K个区间
	for i := K-1; i < ll; i++ {
		tmp := exec(dp, accu, K-1, i) + float64((*accu)[ll-1] - (*accu)[i-1])/float64(ll-i)
		(*dp)[K][ll-1] = fmax((*dp)[K][ll-1], tmp)
	}
	return (*dp)[K][ll-1]
}

func calc(a *[]int) *[]int {
	accu := make([]int, len(*a))
	accu[0] = (*a)[0]
	for i := 1; i < len(*a); i++ {
		accu[i] += accu[i-1] + (*a)[i]
	}
	return &accu
}

func imax(a, b int) int {
	if a < b { return b }
	return a
}

func fmax(a, b float64) float64 {
	if a < b { return b }
	return a
}