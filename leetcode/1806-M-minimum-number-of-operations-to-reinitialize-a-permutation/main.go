package main

import "fmt"

func main() {
	reinitializePermutation(6)
}

func reinitializePermutation(n int) int {
	res := 1

	src_ind := 1
	tar_ind := way(src_ind, n)

	for src_ind != tar_ind {
		fmt.Println(tar_ind)
		tar_ind = way(tar_ind, n)
		res++
	}
	fmt.Println(res)
	return res
}

func way(i int, l int) int {
	if i % 2 == 0 {
		return i / 2
	}
	return l / 2 + (i - 1) / 2
}