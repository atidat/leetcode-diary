package main

import "fmt"

func main() {

}

// 1 <= k <= 10^9
func findMinFibonacciNumbers(k int) int {
	fin := make([]int, 0)
	fin = append(fin, 1, 1)
	prev, tmp := 1, 2
	ind := 2
	for tmp <= 2*k {
		fin = append(fin, tmp)
		prev, tmp = tmp, tmp+prev
		ind++
	}
	fmt.Println(fin)


	var res int
	for 


	return res
}
