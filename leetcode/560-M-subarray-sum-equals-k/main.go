package main

func main() {

}

func subarraySum(nums []int, k int) int {
	var res int

	nl := len(nums)
	sum := make([]int, nl+1, nl+1)
	for i := 1; i <= nl; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	for i := 0; i <= nl; i++ {
		for j := i+1; j <= nl; j++ {
			if sum[j] - sum[i] == k {
				res++
			}
		}
	}
	return res
}