package main

import (
	"fmt"
	"math"
)

func main() {

}

func stoneGameIII(stoneValue []int) string {
	dp := cache(len(stoneValue))

	res := executor(&stoneValue, dp, 0)
	fmt.Println(res)
	if res > 0 {
		return "Alice"
	} else if res < 0 {
		return "Bob"
	}
	return "Tie"
}

func executor(s, dp *[]int, h int) int {
	if h >= len(*s) { return 0 }
	if (*dp)[h] != math.MinInt32 { return (*dp)[h] }

	var plus int
	res := math.MinInt32
	for i := 0; i < 3 && h+i < len(*s); i++ {
		plus += (*s)[h+i]
		res = mmax(res, plus - executor(s, dp, h+1+i))
	}
	(*dp)[h] = mmax(res, (*dp)[h])
	return res
}

func cache(svl int) *[]int {
	dp := make([]int, svl)
	for i := range dp {
		dp[i] = math.MinInt32

	}
	return &dp
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}