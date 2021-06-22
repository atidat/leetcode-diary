package main

import "fmt"

func main() {

}

func totalHammingDistance(nums []int) int {
	dp := make([]int, len(nums))
	dp[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums)-2; i >= 0; i-- {
		dp[i] = nums[i] ^ dp[i+1]
	}
	fmt.Println(dp)
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		res ^= dp[i]
	}
	fmt.Println(res)
	return res
}

func calcHD(a, b int) int {
	c := a ^ b
	var res int
	for c != 0 {
		if c % 2 == 1 {
			res += 1
		}
		c >>= 1
	}
	//fmt.Println(a, b, res)
	return res
}

/*func totalHammingDistance(nums []int) int {
	var res int
	for i := 0; i < len(nums)-1; i++ {
		for j := i+1; j < len(nums); j++ {
			res += calcHD(nums[i], nums[j])
		}
	}
	fmt.Println(res)
	return res
}

func calcHD(a, b int) int {
	c := a ^ b
	var res int
	for c != 0 {
		if c%2 == 1 {
			res += 1
		}
		c >>= 1
	}
	//fmt.Println(a, b, res)
	return res
}*/