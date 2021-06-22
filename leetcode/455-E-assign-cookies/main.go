package main

import "sort"

func main() {

}

func findContentChildren(stu []int, coo []int) int {
	sort.Ints(stu)
	sort.Ints(coo)

	var eat int
	rest := len(coo)-1
	for i := len(stu)-1; i>= 0; i-- {
		for j := rest; j >= 0; j-- {
			if coo[j] >= stu[i] {
				rest = j-1
				eat++
				break
			}
		}
	}
	return eat
}
