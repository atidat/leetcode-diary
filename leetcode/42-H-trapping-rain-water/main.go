package main

import "fmt"

func main() {

}
/*
func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}*/


func trap(height []int) int {
	// 双指针法
	// 前后指针，前后max
	var res int
	lr, rr := 0, len(height)-1
	var lm, rm int//= height[lr], height[rr]
	for lr < rr {
		lm = mmax(lm, height[lr])
		rm = mmax(rm, height[rr])
		if lm < rm {
			res += lm - height[lr]
			lr++
		} else {
			res += rm - height[rr]
			rr--
		}
	}
	fmt.Println(res)
	return res
}



func mmax(a, b int) int {
	if a < b { return b }
	return a
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}

