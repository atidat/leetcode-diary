package main

import "fmt"

func main() {
	reconstructMatrix(2, 1, []int{1,1,1})
	reconstructMatrix(2, 3, []int{2,2,1,1})
	reconstructMatrix(5, 5, []int{2,1,2,0,1,0,1,2,0,1})
	reconstructMatrix(9, 2, []int{0,1,2,0,0,0,0,0,2,1,2,1,2})
}

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	var sums int
	var num2 int
	for i := range colsum {
		sums += colsum[i]
		if colsum[i] == 2 {
			num2++
		}
	}
	if upper + lower != sums || num2 > min(upper, lower) {
		return [][]int{}
	}

	res := make([][]int, 2)
	for i := range res {
		res[i] = make([]int, len(colsum))
	}

	for i := 0; i < len(colsum); i++ {
		if colsum[i] == 2 {
			res[0][i] = 1
			res[1][i] = 1
			upper--
			lower--
		}
	}

	for i := range colsum {
		if colsum[i] == 0 || colsum[i] == 2 {
			continue
		}
		if upper > 0 {
			res[0][i] = 1
			upper--
		} else if lower > 0 {
			res[1][i] = 1
			lower--
		}
	}

	fmt.Println(res)
	fmt.Println()
	return res
}

func min(a, b int) int {
	if a <  b {
		return a
	}
	return b
}