package main

import "sort"

func main() {

}
/*
* 1 <= stones.length <= 30
* 1 <= stones[i] <= 1000
*/
func lastStoneWeightII(stones []int) int {
	sort.Ints(stones)
	return exec(stones)
}

func exec(stones []int) int {
	sl := len(stones)
	if sl == 1 { return stones[0]}

	sub1, sub2 := make([]int, 0), make([]int, 0)
	sum1, sum2 := 0, 0
	for i := range stones {
		if sum1 > sum2 {
			sub2 = append(sub2,stones[i])
			sum2 += stones[i]
		} else {
			sub1 = append(sub1, stones[i])
			sum1 += stones[i]
		}
	}
	return absMinus(exec(sub1), exec(sub2))
}
/*
func mmin(a, b int) int {
	if a < b { return a }
	return b
}
*/
func absMinus(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}