package main

import (
	"fmt"
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		first, second := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return first+second >= second+first
	})

	if nums[0] == 0 {
		return "0"
	}
	res := ""
	for i := 0; i < len(nums); i++ {
		res += strconv.Itoa(nums[i])
	}
	return res
}

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{9, 3, 34, 5, 32}))
	fmt.Println(largestNumber([]int{121, 12}))
	fmt.Println(largestNumber([]int{12, 121}))
	fmt.Println(largestNumber([]int{1, 1, 1}))
	fmt.Println(largestNumber([]int{0, 0}))
	fmt.Println(largestNumber([]int{1, 0, 0}))
}
