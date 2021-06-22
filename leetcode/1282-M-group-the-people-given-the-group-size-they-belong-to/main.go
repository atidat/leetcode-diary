package main

import "fmt"

func main() {
	fmt.Println(groupThePeople([]int{3,3,3,3,3,1,3}))
	fmt.Println("===========================")
	fmt.Println(groupThePeople([]int{2,1,3,3,3,2}))
}

func groupThePeople(groupSizes []int) [][]int {
	res := make([][]int, 0)
	bucket := make(map[int][]int, 0)
	for k, v := range groupSizes {
		if g, ok := bucket[v]; !ok {
			bucket[v] = []int{k}
		} else {
			if len(g) == v {
				res = append(res, g)
				delete(bucket, v)
				bucket[v] = []int{k}
			} else if len(g) < v {
				bucket[v] = append(bucket[v], k)
			}
		}
	}
	for _, v := range bucket {
		res = append(res, v)
	}
	return res
}