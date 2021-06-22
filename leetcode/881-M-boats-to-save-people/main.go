package main

import (
	"fmt"
	"sort"
)

func main() {

}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	pLen := len(people)
	var res int
	big, sma := pLen-1, 0
	for ; big > sma; {
		if people[big] + people[sma] <= limit {
			sma++
		}
		res++
		big--
	}
	if big == sma {
		res++
	}
	fmt.Println(res)
	return res
}