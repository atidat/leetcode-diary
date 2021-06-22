package main

import "math"

func main() {

}

func accu(arr *[]int) *[]int {
	total := make([]int, len(*arr)+1)
	for i := 1; i <= len(*arr); i++ {
		total[i] = total[i-1] + (*arr)[i-1]
	}
	return &total
}

func cache(sl int) *[][]int {
	dp := make([][]int, sl)
	for i := range dp {
		dp[i] = make([]int, sl)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	return &dp
}

func stoneGameVII(stones []int) int {
	to := accu(&stones)
	dp := cache(len(stones))
	return executor(to, dp, 0, len(stones)-1)
}

func executor(to *[]int, dp *[][]int, h, t int) int {
	if h >= t { return 0 }
	if (*dp)[h][t] != math.MinInt32 { return (*dp)[h][t] }

	res := math.MinInt32
	res = mmax(res, (*to)[t+1] - (*to)[h+1] - executor(to, dp, h+1, t))
	res = mmax(res, (*to)[t] - (*to)[h] - executor(to, dp, h, t-1))

	(*dp)[h][t] = res
	return res
}

func mmax(a, b int) int {
	if a < b {return b}
	return a
}

/*func stoneGameVII(stones []int) int {
	fmt.Println(stones)
	dp := pre(len(stones), &stones)
	total := accu(&stones)
	res := executor(dp, total, 0, len(stones)-1)
	fmt.Println(*dp)
	return res
}

func pre(sl int, s *[]int) *[][]int {
	dp := make([][]int, sl)
	for i := range dp {
		dp[i] = make([]int, sl)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}
	return &dp
}

func accu(s *[]int) *[]int {
	total :=  make([]int, len(*s)+1)
	for i := 1; i <= len(*s); i++ {
		total[i] = total[i-1] + (*s)[i-1]
	}
	return &total
}

func executor(dp *[][]int, to *[]int, h, t int) int {
	if h >= t { return 0 }
	if (*dp)[h][t] != math.MinInt32 {
		return (*dp)[h][t]
	}

	// dp表示(alice-bob)的差值
	res := mmax((*to)[t+1] - (*to)[h+1] - executor(dp, to, h+1, t),
		(*to)[t] - (*to)[h] - executor(dp, to, h, t-1))
	(*dp)[h][t] = res
	return res
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}

func mmin(a, b int) int {
	if a <  b { return a }
	return b
}*/