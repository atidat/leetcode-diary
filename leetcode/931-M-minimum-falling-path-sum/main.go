package main

import (
	"fmt"
	"math"
)

func main() {

}

func minFallingPathSum(A [][]int) int {
	al := len(A)
	for i := 1; i < al; i++ {
		for j := 0; j < al; j++ {
			tmp := A[i-1][j]
			if j-1 >= 0 {
				tmp = mmin(tmp, A[i-1][j-1])
			}
			if j+1 < al {
				tmp = mmin(tmp, A[i-1][j+1])
			}
			A[i][j] += tmp
		}
	}
	fmt.Println(A)

	res := math.MaxInt32
	for i := range A[al-1] {
		res = mmin(res, A[al-1][i])
	}
	return res
}

func mmin(a, b int) int {
	if a < b { return a }
	return b
}
func mmax(a, b int) int {
	if a < b { return b }
	return a
}