package main

import
(
	"fmt"
	"sort"
)

func main() {
	//arr := []int{3,5,1,6,8}
	//arr := []int{3,4,1,2,4}
	//arr := []int{2,6,2,2,7}
	arr := []int{2,14,15,17,6,18,12,18,15,4} //117
	sort.Ints(arr)
	fmt.Printf("\nfinal value is %d", maxSumDivThree(arr))
}

/*
【可被3整除的最大和】
给你一个整数数组 nums，请你找出并返回能被三整除的元素最大和。

*/
func maxSumDivThree(nums []int) int {
	total, oneArr, twoArr := getTotalSumAnd12Arrays(nums)
	if total % 3 == 0 {
		return total
	} else if total % 3 == 1 {
		return perMax(handle12Array(total, oneArr, true), handle12Array(total, twoArr, false))
	}

	return perMax(handle12Array(total, oneArr, false), handle12Array(total, twoArr, true))
}

func getTotalSumAnd12Arrays(nums []int) (int, []int, []int) {
	sum := 0
	oneArr, twoArr := make([]int, 0), make([]int, 0)
	for _, v := range nums {
		sum += v

		if v % 3 == 1 {
			oneArr = append(oneArr, v)
		} else if v % 3 == 2 {
			twoArr = append(twoArr, v)
		}
	}
	fmt.Printf("sum: %d, oneArr: %v, twoArr: %v", sum, oneArr, twoArr)
	return sum, oneArr, twoArr
}

func handle12Array(total int, arr []int, status bool) int {
	// status == true, 减去1个值
	// status == false, 减去2个值
	if status && len(arr) < 1 || !status && len(arr) < 2 {
		return 0
	}

	sort.Ints(arr)
	if status {
		return total - arr[0]
	}
	return total - arr[0] - arr[1]
}

func perMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}