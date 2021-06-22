package main

func getMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func origin(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]

	for k := 1; k < len(nums); k++ {
		dp[k+1] = getMax(dp[k-1]+nums[k], dp[k])
	}

	return dp[len(nums)]
}

func rob(nums []int) int {
	//	fmt.Println(nums, " ======")
	lenNums := len(nums)
	switch lenNums {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return getMax(nums[0], nums[1])
	}

	withoutTail := origin(nums[:lenNums-1])
	withoutHead := origin(nums[1:])
	//withBody := origin(nums[1 : lenNums-1])

	//return getMax(getMax(withoutHead, withoutTail), withBody)
	return getMax(withoutHead, withoutTail)
}
