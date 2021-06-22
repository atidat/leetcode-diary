package main

func majorityElement(nums []int) int {
	lenN := len(nums)
	halfL := lenN / 2
	repo := make(map[int]int, 0)
	for i := 0; i < lenN; i++ {
		time, ok := repo[nums[i]]
		if ok {
			time++
			if time > halfL {
				return nums[i]
			}
		} else {
			repo[nums[i]] = 1
		}
	}
	return 0
}

func main() {

}
