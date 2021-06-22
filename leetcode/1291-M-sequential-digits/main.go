package main

import "fmt"

func main() {
	sequentialDigits(10,1000000000)
}

func sequentialDigits(low int, high int) []int {
	start := []int{12, 123, 1234, 12345, 123456, 1234567, 12345678, 123456789}
	incr := []int{11, 111, 1111, 11111, 111111, 1111111, 11111111, 111111111}
	var ind int
	for i := ind; i < 8; i++ {
		ind = i
		if low <= start[i] {
			break
		}
	}
	var val = start[ind-1]
	for val > low {
		val += incr[ind-1]
	}

	res := make([]int, 0)
	for val <= high {
		res = append(res, val)
		if val % 10 == 9 {
			ind++
			if ind >= 8 {
				break
			}
			val = start[ind]
		} else {
			val += incr[ind]
		}
	}
	fmt.Println(res)
	return res
}