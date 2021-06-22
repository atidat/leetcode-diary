package main

import (
	"fmt"
	"sort"
)

func main() {

}

func subsetXORSum(nums []int) int {
	sort.Ints(nums)

	var sum int
	nl := len(nums)
	for i := 1; i <= nl; i++ {
		dfs(nums,i, 0, 0, &sum)
	}
	fmt.Println(sum)
	return sum
}

func dfs(nums []int, n int, ind int, res int, sum *int) {
	if n == 0 {
		*sum += res
		return
	}

	for i := ind; i < len(nums); i++ {
		dfs(nums, n-1, i+1, res^nums[i], sum)
	}
}