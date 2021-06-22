package main

import "fmt"

func getMin(nums []int) int {
	minNum := nums[0]
	for k := range nums {
		if minNum > nums[k] {
			minNum = nums[k]
		}
	}
	return minNum
}

func minNum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(tris [][]int) int {
	rows := len(tris)
	if rows == 1 {
		return getMin(tris[0])
	}

	for k := rows - 2; k >= 0; k-- {
		for i := 0; i <= k; i++ {
			tris[k][i] = minNum(tris[k][i]+tris[k+1][i], tris[k][i]+tris[k+1][i+1])
		}
		fmt.Println(tris[k])
	}
	return tris[0][0]
}

func main() {
}
