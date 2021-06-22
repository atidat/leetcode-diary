package main

import (
	"fmt"
	"sort"
)

func maxProfit(k int, prices []int) int {
	profits := make([]int, 0)
	for i := 0; i < len(prices)-1; {
		ind, pro := shortterm(prices, i)
		i = ind
		if pro != -1 {
			profits = append(profits, pro)
			i++
		}
	}

	profits = sort.IntSlice(profits)
	res, picks := 0, k
	if picks > len(profits) {
		picks = len(profits)
	}
	fmt.Println(profits)
	for i := 0; i < picks; i++ {
		res += profits[i]
	}
	return res
}

func shortterm(prices []int, head int) (int, int) {
	max, min := prices[head], prices[head]
	rcd := -1
	for i := head + 1; i < len(prices); i++ {
		if prices[i] <= min {
			if rcd == -1 {
				return i, -1
			} else {
				break
			}
		}
		if prices[i] > max {
			rcd = i
			max = prices[i]
		}
	}
	return rcd, prices[rcd] - prices[head]
}

func main() {
	//fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
	//fmt.Println(maxProfit(2, []int{2, 4, 1}))
	//fmt.Println(maxProfit(1, []int{3, 3}))
	fmt.Println(maxProfit(2, []int{6, 1, 3, 2, 4, 7}))
}
