package main

import "math"

func main() {

}

func stoneGameV(stoneValue []int) int {
	dp := cache(len(stoneValue))
	to := accu(&stoneValue)
	return executor(to, dp,0, len(stoneValue)-1)
}

func cache(svl int) *[][]int {
	dp := make([][]int, svl)
	for i := range dp {
		dp[i] = make([]int, svl)
	}
	return &dp
}

func accu(s *[]int) *[]int {
	total := make([]int, len(*s)+1)
	for i := 1; i <= len(*s); i++ {
		total[i] = total[i-1] + (*s)[i-1]
	}
	return &total
}

func executor(to *[]int, dp *[][]int, h, t int) int {
	if h >= t { return 0 }
	if (*dp)[h][t] != 0 { return (*dp)[h][t] }
	/*
	* dp[h][t] = max(
	*	suml + executor(h, i), if suml < sumr
	*	sumr + executor(i+1, t), if suml > sumr
	*	suml + max(executor(h, i) executor(i+1, t)), if suml = sumr
	* )
	*/
	res := math.MinInt32
	for i := h; i <= t; i++ {
		suml := (*to)[i+1] - (*to)[h]
		sumr := (*to)[t+1] - (*to)[i+1]

		if suml < sumr {
			res = mmax(res, suml + executor(to, dp, h, i))
		} else if suml > sumr {
			res = mmax(res, sumr + executor(to, dp, i+1, t))
		} else {
			res = mmax(res, suml+mmax(executor(to, dp, h, i), executor(to, dp, i+1, t)))
		}
	}
	(*dp)[h][t] = res
	return res
}

func mmax(a, b int) int {
	if a < b {return b}
	return a
}

/*func stoneGameV(stoneValue []int) int {
	accu := accusum(&stoneValue)
	dp := cache(len(stoneValue))
	return executor(accu, dp, 0, len(stoneValue)-1)
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}

func mmax(a, b int) int {
	return a + b - mmin(a, b)
}

func accusum(sv *[]int) *[]int {
	accu := make([]int, len(*sv)+1)
	for i := 1; i < len(*sv)+1; i++ {
		accu[i] = accu[i-1] + (*sv)[i-1]
	}
	return &accu
}

func cache(svl int) *[][]int {
	dp := make([][]int, svl)
	for i := range dp {
		dp[i] = make([]int, svl)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return &dp
}

func executor(a *[]int, dp *[][]int, h, t int) int {
	if h >= t { return 0 }
	if (*dp)[h][t] != -1 { return (*dp)[h][t] }

	var res int
	for i := h; i <= t; i++ {
		suml := (*a)[i+1] - (*a)[h]
		sumr := (*a)[t+1] - (*a)[i+1]

		if suml < sumr {
			res = mmax(res, suml + executor(a, dp, h, i))
		} else if suml > sumr {
			res = mmax(res, sumr + executor(a, dp, i+1, t))
		} else {
			res = mmax(res, suml + mmax(executor(a, dp, h, i), executor(a, dp, i+1, t)))
		}
	}
	(*dp)[h][t] = res
	return res
}*/