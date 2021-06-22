package main

import (
	"fmt"
	"sort"
)

func main() {

}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	// solution 1
/*	var res int
	for i := 0; i < len(houses); i++ {
		j := 0
		for ; j+1 < len(heaters); j++ {
			if abs(heaters[j] - houses[i]) < abs(heaters[j+1]-houses[i]) {
				break
			}
		}
		res = mmax(res, abs(heaters[j] - houses[i]))
	}
	fmt.Println(res)
	return res*/

	// solution 2

}

func abs(a int) int {
	if a < 0 { return 0-a }
	return a
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}