package main

func main() {

}

func search(nums []int, target int) int {

	if len(nums) == 1 {
		if nums[0] == target { return 0	}
		return -1
	}
	if len(nums) == 2 {
		if nums[0] == target { return 0 }
		if nums[1] == target { return 1 }
		return -1
	}

	res := -1
	bot := findBot(nums)
	if target >= nums[0] {
		res = erfen(nums[:bot], target)
	} else if target <= nums[len(nums)-1] {
		res = bot
		tmp := erfen(nums[bot:], target)
		if tmp == -1 {
			res = -1
		} else {
			res += tmp
		}
	}
	return res
}


func findBot(nums []int) int {
	h, t := 0, len(nums)-1
	m := (h+t)/2
	for h <= m {
		v1, v2 := nums[m]-nums[m-1], nums[m]-nums[m+1]
		if v1 * v2 > 0 {
			if v1 < 0 {
				return m
			}
			return m+1
		}

		if nums[m] >= nums[h] && nums[m] >= nums[t] {
			h = m
		} else if nums[m] <= nums[h] && nums[m] <= nums[t] {
			t = m
		} else {
			return len(nums)
		}
		m = (h+t)/2
	}
	return m
}

func erfen(arr []int, target int) int {
	h, t := 0, len(arr)-1
	m := (h+t)/2
	for h <= t {
		if arr[m] == target {
			return m
		}
		if arr[m] > target {
			t = m-1
		} else {
			h = m+1
		}
		m = (h+t)/2
	}
	return -1
}