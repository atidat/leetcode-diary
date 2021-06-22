package main

import (
	"math"
)

func maxProduct(nums []int) int {

	// minNum := math.MaxInt32
	// maxNum := math.MinInt32
	minNum, maxNum := 1, 1
	hist := math.MinInt32

	for k := range nums {
		tmp := minNum
		minNum = threeMin(minNum*nums[k], maxNum*nums[k], nums[k])
		maxNum = threeMax(maxNum*nums[k], tmp*nums[k], nums[k])
		hist = threeMax(hist, minNum, maxNum)

	}

	return hist
}

func threeMax(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if a <= b {
		if b > c {
			return b
		}
	}
	return c
}

func threeMin(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if a >= b {
		if b > c {
			return c
		}
	}
	return b
}

func main() {

}
