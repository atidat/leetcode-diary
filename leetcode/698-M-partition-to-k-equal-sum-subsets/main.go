package main

import "sort"

func main() {

}

func canPartitionKSubsets(nums []int, k int) bool {
	m := make(map[int][]int)
	var sum int
	for i := range nums {
		if v, ok := m[nums[i]]; !ok {
			m[nums[i]] = []int{i}
		} else {
			m[nums[i]] = append(v, i)
		}
		sum += nums[i]
	}
	if sum % k != 0 { return false }
	avg := sum / k

	sort.Ints(nums)
	for i := len(nums)-1; i >= 0; i-- {
		if _, ok := m[nums[i]]; !ok {
			continue
		}

	}
}