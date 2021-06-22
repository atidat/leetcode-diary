package main

import "fmt"

func main() {
}

/*
Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k]
such that i < j < k and nums[i] < nums[k] < nums[j].

Return true if there is a 132 pattern in nums, otherwise return false.

Constraints:
	n == nums.length
	1 <= n <= 3 * 10^4
	-10^9 <= nums[i] <= 10^9

*/
/*
func find132pattern(nums []int) bool {
	fmt.Println(nums)
	if len(nums) < 3 {
		return false
	}

	index := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			index = append(index, i+1)
			continue
		}

		if len(index) == 0 {
			continue
		}
		fmt.Println(index)
		anchor := index[0]-1
		for j := 0; j < len(index); j++ {
			for k := 0; k < anchor; k++ {
				if nums[k] < nums[index[j]] {
					return true
				}
			}
		}
		index = []int{}
	}

	fmt.Println(index)
	if len(index) == 0 {
		return false
	}

	anchor := index[0]-1
	for j := 0; j < len(index); j++ {
		for k := 0; k < anchor; k++ {
			if nums[k] < nums[index[j]] {
				return true
			}
		}
	}
	return false
}
*/
// 栈
func find132pattern(nums []int) bool {
	fmt.Println(nums)
	if len(nums) < 3 {
		return false
	}

	for i := 1; i < len(nums); i++ {
		down := fstDownPoint(nums, i)
		if down == -1 {
			return false
		}
		bigger := fstBiggerPoint(nums, down+1)

		// fmt.Println(down, bigger)

		for j := 0; j < down; j++ {
			for k := down+1; k <= bigger; k++ {
				if nums[j] < nums[k] && nums[k] < nums[down]{
					return true
				}
			}
		}
	}
	return false
}

func fstDownPoint(nums []int, start int) int {
	for i := start; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}

func fstBiggerPoint(nums []int, start int) int {
	for i := start; i < len(nums); i++ {
		if nums[i] >= nums[start-1] {
			return i
		}
	}
	return len(nums)-1
}