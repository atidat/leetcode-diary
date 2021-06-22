package main

/*import (
	"fmt"
	"time"
)*/

func main() {

}

func findLength(A []int, B []int) int {
	// return dyPro(A, B)
	if len(A) < len(B) {
		return rollArr(A, B)
	}
	return rollArr(B, A)
}
/*func dyPro(A, B []int) int {
	// dp[i][j] means 包含A[i] & B[j]的最长公共子数组长度
	// dp[i][j] = 0, A[i] != B[j]
	// dp[i][j] = dp[i-1][j-1] + 1, A[i] == B[j]
	al, bl := len(A), len(B)
	dp := make([][]int, bl)
	for i := range dp {
		dp[i] = make([]int, al)
	}
	for j := 0; j < al; j++ {
		if B[0] == A[j] {
			dp[0][j] = 1
		}
	}
	for i := 0; i < bl; i++ {
		if A[0] == B[i] {
			dp[i][0] = 1
		}
	}
	var sum int
	for j := 1; j < bl; j++ {
		for i := 1; i < al; i++ {
			if A[i] == B[j] {
				dp[j][i] = dp[j-1][i-1] + 1
				sum = mmax(sum, dp[j][i])
			}
		}
	}
	//fmt.Println(sum)
	return sum
}*/

func mmax(a, b int) int {
	if a < b {return b}
	return a
}
func mmin(a, b int) int {
	if a < b {return a}
	return b
}

func rollArr(A, B []int) int {
	// A是短的数组，B是长的数组
	al, bl := len(A), len(B)
	var sum, tmp int
	for i := 1; i < al+bl-1; i++ {

		var fh, sh, tarl int
		if i < al {
			fh = al-i
			sh = 0
			tarl = i
		} else {
			fh = 0
			sh = i - al
			tarl = mmin(al, bl-sh)
		}

		tmp = calc(&A, &B, fh, sh, tarl)
		sum = mmax(sum, tmp)
	}
	return sum
}

func calc(arrf, arrs *[]int, fh, sh, tarl int) int {
	var sum, tmp int
	for i := 0; i < tarl; i++ {
		if (*arrf)[fh+i] == (*arrs)[sh+i] {
			tmp++
			sum = mmax(sum, tmp)
		} else {
			tmp = 0
		}
	}
	return sum
}

