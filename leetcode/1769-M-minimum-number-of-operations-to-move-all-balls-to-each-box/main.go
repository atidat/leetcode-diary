package main

import "fmt"

func main() {
	minOperations("110")
	minOperations("001011")
}

func minOperations(boxes string) []int {
	ones := make(map[int]bool, 0)
	for i := range boxes {
		if boxes[i] == '1' {
			ones[i] = true
		}
	}

	res := make([]int, len(boxes))
	for i := range boxes {
		tmp := 0
		for k, _ := range ones {
			tmp += abs(k, i)
		}
		res[i] = tmp
	}
	fmt.Println(res)
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}