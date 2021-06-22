package main

import "fmt"

func main() {

}

func partitionLabels(s string) []int {
	bucket := make(map[int32]int, 0)
	for k, v := range s {
		if _, ok := bucket[v]; !ok {
			bucket[v] = k
		} else {
			bucket[v] = k
		}
	}
	fmt.Println(bucket)

	start, last := 0, -1
	res := make([]int, 0)
	for k, v := range s {

		tmp, _ := bucket[v]
		last = mmax(last, tmp)
		if k == last {
			res = append(res, last - start + 1)
			start = last + 1
		}
	}
	fmt.Println(res)
	return res
}

func mmax(a, b int) int {
	if a >  b {
		return a
	}
	return b
}