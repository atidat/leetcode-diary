package main

import "fmt"

func main() {
}

/*
	Write a program to find the n-th ugly number.
	Ugly numbers are positive integers which are divisible by a or b or c.

	Constraints:

		1 <= n, a, b, c <= 10^9
		1 <= a * b * c <= 10^18
		It's guaranteed that the result will be in range [1, 2 * 10^9]
*/
func nthUglyNumber(n int, a int, b int, c int) int {
	fmt.Println("n, a, b,c ", n,a,b,c)
	var val int
	if a > b {
		if a > c {
			val = a
		} else {
			val = c
		}
	} else {
		if b > c {
			val = b
		} else {
			val = c
		}
	}

	left, right := 1, val * n
	var mid int
	for left < right {
		mid = left + (right - left)/2
		nums := mid/a + mid/b + mid/c - mid/leastComMul(a,b) /
		- mid/leastComMul(a,c) - mid/leastComMul(b,c) /
		+ mid/leastComMul(a, leastComMul(b,c))
		fmt.Println(left, mid, right, nums)
		if nums < n {
			left = mid + 1
		} else {
			right = mid
		}
	}
	fmt.Println(left, right)
	return left
}

// 最小公倍数
func leastComMul(a, b int) int {
	return (a*b)/greatestComDiv(a, b)
}

// 最大公约数
func greatestComDiv(a, b int) int {
	if a == 0 {
		return b
	}
	return greatestComDiv(b%a, a)
}