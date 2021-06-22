package main

func main() {

}

/*func checkSubarraySum(nums []int, k int) bool {
	if k < 0 { k = 0 - k }

	var sum int
	m := make(map[int]int)
	m[0] = -1
	for i := range nums {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}

		if ind, ok := m[sum]; ok {
			if i-ind > 1 { return true}
		} else {
			m[sum] = i
		}
	}
	return false
}
*/

func checkSubarraySum(nums []int, k int) bool {
	sum := make([]int, 1+len(nums))
	for i := 0; i < len(nums); i++ {
		sum[i+1] = sum[i] + nums[i]
		for j := 0; j < i; j++ {
			if (sum[i+1] - sum[j]) % k == 0 {
				return true
			}
		}
	}
	return false
}