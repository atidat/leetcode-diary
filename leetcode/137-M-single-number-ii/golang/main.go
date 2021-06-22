package main

func main() {

}

func singleNumber(nums []int) int {
	fstSus := nums[0]
	var sndSus int
	var ind int
	for ind = 1; ind < len(nums); ind++ {
		if fstSus != nums[ind] {
			sndSus = nums[ind]
			break
		}
	}

	for ; ind < len(nums); ind++ {
		if fstSus == nums[ind] {
			return sndSus
		}
	}
	return fstSus
}
