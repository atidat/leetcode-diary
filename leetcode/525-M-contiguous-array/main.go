package main

import "fmt"

func main() {

}

func findMaxLength(nums []int) int {
	nl := len(nums)

	f := make([][]int, nl)
	for i := 0; i < nl; i++ {
		f[i] = make([]int, nl)
		if i+1 < nl && nums[i] != nums[i+1] {
			f[i][i+1] = 2
		}
	}

	prev := make([]int, nl)
	cur := make([]int, nl)
	for i := nl-2; i >= 0; i-- {

		var curSum = 0
		var one, zero int
		if nums[i] != nums[i+1] {
			one, zero = 1, 1
			curSum = 2
		} else if nums[i] == 0 {
			zero = 2
		} else {
			one = 2
		}

		for j := i+2; j < nl; j++ {
			if nums[j] == 0 {
				zero++
			} else {
				one++
			}
			if zero == one {
				curSum = 2 * one
			}
			cur[i][j] = mmax(curSum, mmax(cur[i+1][j], f[i][j-1]))
		}
	}
	fmt.Println(f[0][nl-1])
	return f[0][nl-1]
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}