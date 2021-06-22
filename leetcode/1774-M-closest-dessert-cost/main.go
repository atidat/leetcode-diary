package main

import "fmt"

func main() {

}
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	/*
	n == baseCosts.length
	m == toppingCosts.length
	1 <= n, m <= 10
	1 <= baseCosts[i], toppingCosts[i] <= 104
	1 <= target <= 104
	*/
	who := make([]bool, 20000)

	for _, v := range baseCosts {
		who[v] = true
	}


	for i := 1; i < 20000; i++ {

		for _, v := range toppingCosts {
			if i-v >= 0 {
				who[i] = who[i] || who[i-v]
			}
		}
	}
	fmt.Println(who[0:20])
	for i := 1; i < 20000; i++ {
		for _, v := range toppingCosts {
			if i-v >= 0 {
				who[i] = who[i] || who[i-v]
			}
		}
	}
	fmt.Println(who[0:20])
	/*for _, v := range toppingCosts {
		for i := 19999; i >= v; i-- {
			who[i] = who[i] || who[i-v]
		}
	}
	fmt.Println(who[0:20])
	for _, v := range toppingCosts {
		for i := 19999; i >= v; i-- {
			who[i] = who[i] || who[i-v]
		}
	}
	fmt.Println(who[0:20])*/

	var res int
	for i := 0; target-i >= 0 || target+i < 20000; i++ {
		if target-i >= 0 && who[target-i] {
			fmt.Println(target-i)
			return target-i
		}
		if target+i < 20000 && who[target+i] {
			fmt.Println(target+i)
			return target+i
		}
	}
	fmt.Println(res)
	return res
}