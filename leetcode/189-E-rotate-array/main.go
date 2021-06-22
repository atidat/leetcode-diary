package main

import "fmt"

func rotate(nums []int, k int) {
	lenN := len(nums)
	if lenN == 0 {
		return
	}

	offset := k % lenN
	nums = append(nums[lenN-offset:], nums[:lenN-offset]...)
	fmt.Println(nums)
	return
}
func main() {
}
