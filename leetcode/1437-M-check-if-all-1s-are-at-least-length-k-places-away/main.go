package main

func main() {

}

func kLengthApart(nums []int, k int) bool {
	if k == 0 {
		return true
	}

	for i := 0; i < len(nums); {
		if !smallLoop(nums, i, k) {
			return false
		}
		for j := i + 1 + k; j <= len(nums); j++ {
			if j >= len(nums) {
				return true
			}
			if nums[j] == 1 {
				i = j
				break
			}
		}
	}
	return true
}

func smallLoop(nums []int, index, scope int) bool{
	for i := 1; i <= scope; i++ {
		if nums[index + i] == 1 {
			return false
		}
	}
	return true
}