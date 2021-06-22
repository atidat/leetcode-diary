package main

import "fmt"

func main() {
	mat := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	k1, k2 := 1, 2
	res1 := matrixBlockSum(mat, k1)
	res2 := matrixBlockSum(mat, k2)
	printDebug(res1)
	printDebug(res2)
}

func printDebug(res [][]int) {
	for i := range res {
		fmt.Println(res[i])
	}
	fmt.Println("==========")
}

func matrixBlockSum(mat [][]int, K int) [][]int {
	
}