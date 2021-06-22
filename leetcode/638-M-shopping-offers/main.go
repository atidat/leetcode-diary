package main

import (
	"fmt"
)

func main() {

}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	res := executor(&price, &special, &needs)
	fmt.Println(res)
	return res
}

func executor(price *[]int, special *[][]int, needs *[]int) int {
	res := 0
	for i := range *needs {
		res += (*needs)[i] * (*price)[i]
	}

	for i := range *special {
		// needs不满足当前套餐
		if invalidToBuy(needs, &(*special)[i]) {
			continue
		}
		buyProd(needs, &(*special)[i])
		res = mmin(res, executor(price, special, needs)+
			(*special)[i][len(*needs)])
		buyRecover(needs, &(*special)[i])

	}
	return res
	//return res+buyRest(needs, price)
}

func mmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func invalidToBuy(needs, curProds *[]int) bool {
	for i := range *needs {
		if (*needs)[i] - (*curProds)[i] < 0 {
			return true
		}
	}
	return false
}

func buyProd(needs, curProds *[]int) {
	for i := range *needs {	(*needs)[i] -= (*curProds)[i] }
}

func buyRecover(needs, curProds *[]int) {
	for i := range *needs {	(*needs)[i] += (*curProds)[i] }
}

func buyRest(needs, prices *[]int) int {
	res := 0
	for i := range *needs {	res += (*needs)[i] * (*prices)[i] }
	return res
}