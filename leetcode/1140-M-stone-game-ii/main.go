package main

import (
	"fmt"
	"math"
)

func main() {

}


func stoneGameII(piles []int) int {
	dp := cache(len(piles))
	res := executor(&piles, dp, 0, 1)
	sum := total(&piles, 0, len(piles))
	fmt.Println((res+sum)/2)
	return (res+sum)/2
}

func total(p *[]int, h, t int) int {
	var res int
	for i := h; i < t; i++ {
		res += (*p)[i]
	}
	return res
}

func executor(p, dp *[]int, h, M int) int {
	if h+2*M >= len(*p) {
		res := total(p, h, len(*p))
		(*dp)[h] = mmax((*dp)[h], res)
		return (*dp)[h]

	}

	if (*dp)[h] != math.MinInt32 {
		return (*dp)[h]
	}
	var plus int
	res := math.MinInt32
	for i := 1; i <= 2*M && h+i-1 < len(*p); i++ {
		plus += (*p)[h+i-1]
		res = mmax(res, plus - executor(p, dp, h+i, mmax(i, M)))
	}
	(*dp)[h] = mmax((*dp)[h], res)
	return res
}

func cache(pl int) *[]int {
	dp := make([]int, pl)
	for i := range dp {
		dp[i] = math.MinInt32
	}
	return &dp
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}

	/*dp := cache(&piles)
	res := executor(&piles, dp,0, 1)
	return (total(&piles, 0, len(piles)) + res)/2
}

func executor(p *[]int, dp *[][]int, h, M int) int {
	if h+2*M >= len(*p) {
		return total(p, h, len(*p))
	}

	if (*dp)[h][M] != -1 { return (*dp)[h][M] }

	var accu int
	res := math.MinInt32
	for i := 1; i <= 2*M ; i++ {
		accu += (*p)[h+i-1]
		res = mmax(res, accu - executor(p, dp, h+i, mmax(i, M)))
	}
	(*dp)[h][M] = res
	return res
}

func total(p *[]int, h, t int) int {
	var res int
	for i := h; i < t; i++ {
		res += (*p)[i]
	}
	return res
}

func cache(p *[]int) *[][]int {
	pl := len(*p)
	dp := make([][]int, pl)
	for i := range dp {
		dp[i] = make([]int, pl+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return &dp
}
*/
