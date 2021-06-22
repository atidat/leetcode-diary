package main

import "fmt"

func main() {
	getSumAbsoluteDifferences([]int{2,3,5})
	getSumAbsoluteDifferences([]int{1,4,6,8,10})
}

func getSumAbsoluteDifferences(nums []int) []int {
	ls, rs := sum(&nums)
	res := make([]int, len(nums))

	ln := len(nums)-1

	for i := 0; i < ln; i++ {
		res[i] = calcLeft(nums[i], i, ls) + calcRight(nums[i], ln, i, rs)
	}
	res[ln] = (ln+1)*nums[ln] - (*ls)[ln]
	fmt.Println(res)
	return res
}

func sum(nums *[]int) (*[]int, *[]int) {
	ls := make([]int, len(*nums))
	ls[0] = (*nums)[0]

	rs := make([]int, len(*nums))
	ind := len(*nums)-1
	rs[ind] = (*nums)[ind]

	for i := 1; i < len(*nums); i++ {
		ls[i] = ls[i-1] +(*nums)[i]
		rs[ind-i] = rs[ind-i+1] + (*nums)[ind-i]
	}
	return &ls, &rs
}

func calcLeft(cur int, n int, ls *[]int) int {
	return diffAbs(cur*(n+1), (*ls)[n])
}

func calcRight(cur int, l, n int, rs *[]int) int {
	return diffAbs(cur*(l-n), (*rs)[n+1])
}

func diffAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}