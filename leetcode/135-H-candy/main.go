package main

import "fmt"

func main() {

}

func candy(ratings []int) int {
	rLen := len(ratings)
	if rLen < 2 { return rLen }

	dp := make([]int, rLen)
	dp[0] = 1
	for i := 1; i < rLen; i++ {
		fmt.Println(dp)
		if ratings[i-1] < ratings[i] {
			dp[i] = dp[i-1] + 1
			continue
		}
		if dp[i-1] == 1 && ratings[i-1] != ratings[i]{
			influence(&dp, i, &ratings)
		} else {
			dp[i] = 1
		}
	}
	return calcSum(&dp)
}

func calcSum(dp *[]int) int {
	fmt.Println(*dp)
	sum := 0
	for i := range (*dp) {
		sum += (*dp)[i]
	}
	return sum
}

func influence(dp *[]int, i int, r *[]int) {
	(*dp)[i] = 1
	for j := i-1; j >= 0; j-- {
		(*dp)[j] = (*dp)[j+1] + 1

		if j == 0 {
			break
		}
		if (*r)[j-1] > (*r)[j] && (*dp)[j-1] <= (*dp)[j] {
			continue
		} else {
			break
		}
	}
}