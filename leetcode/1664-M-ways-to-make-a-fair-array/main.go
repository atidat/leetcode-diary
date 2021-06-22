package main

import "fmt"

func main() {

}

func waysToMakeFair(nums []int) int {
	// odd:奇数  even:偶数
	var odd_sum, even_sum int
	for i := range nums {
		if i % 2 == 0 {
			even_sum +=nums[i]
		} else {
			odd_sum += nums[i]
		}
	}

	var res int
	var left_odd, left_even, right_odd, right_even int
	for i := range nums {
		if i % 2 != 0 {
			left_odd += nums[i] //
		} else {
			left_even += nums[i]
		}

		right_even = odd_sum - left_odd //
		right_odd = even_sum - left_even

		if i % 2 != 0 {
			right_odd -= nums[i] //
		} else {
			right_even -= nums[i]
		}

		if left_odd + right_odd == left_even + right_even {
			res += 1
		}
	}
	fmt.Println(res)
	return res
}