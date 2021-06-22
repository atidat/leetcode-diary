package main

func main() {

}

func twoSum(numbers []int, target int) []int {
	if len(numbers) == 0 || numbers[0] >= target {
		return []int{}
	}

	for head, tail := 0, len(numbers); head < tail; {
		if numbers[head]+numbers[tail] == target {
			return []int{head, tail}
		}

		if numbers[head]+numbers[tail] < target {
			head++
		} else {
			tail--
		}
	}
	return []int{}
}
