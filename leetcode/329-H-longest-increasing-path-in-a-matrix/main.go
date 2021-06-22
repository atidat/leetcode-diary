package main

import (
	"fmt"
	"math"
)

func main() {

}

// 矩阵中的最长递增路径
func longestIncreasingPath(matrix [][]int) int {
	repoTmp := calc1(matrix)
	return calc2(matrix, repoTmp)
}

func calc1(r [][]int) [][]int {
	repoTmp := make([][]int, len(r))
	for i := 0; i < len(r); i++ {
		repoTmp[i] = make([]int, len(r[0]))
		for j := 0; j < len(r[0]); j++ {
			repoTmp[i][j] = 1
		}
	}
	for i := 0; i < len(r); i++ {
		for j := 0; j < len(r[0]); j++ {
			if i-1 >= 0 && r[i][j] > r[i-1][j] {
				repoTmp[i][j] = mmax(repoTmp[i][j], repoTmp[i-1][j] + 1)
			}
			if j-1 >= 0 && r[i][j] > r[i][j-1] {
				repoTmp[i][j] = mmax(repoTmp[i][j], repoTmp[i][j-1] + 1)
			}

			if i+1 < len(r) && r[i][j] > r[i+1][j] {
				repoTmp[i][j] = mmax(repoTmp[i][j], repoTmp[i+1][j] + 1)
			}
			if j+1 < len(r[0]) && r[i][j] > r[i][j+1] {
				repoTmp[i][j] = mmax(repoTmp[i][j], repoTmp[i][j+1] + 1)
			}
		}
	}
	fmt.Printf("up-left -> down-right: %+v\n", repoTmp, )
	return repoTmp
}

// 从右下角往左斜上递归
func calc2(r [][]int, repoTmp [][]int) int {
	var res = math.MinInt32
	// 每个点c的值取决与c下面的值和c右边的值
	// 如果 c值大于 c下和c右，则c递增 = max(c下递增， c右递增) + 1
	// 如果 c值小于 c下和c右，则c递增 = 1
	// 如果 c值大于min(c下, c右), 小于 max(c下, c右)，则c递增 = min(c下, c右) +

	for i := len(r)-1; i >= 0; i-- {
		for j := len(r[0])-1; j >= 0; j-- {
			if i-1 >= 0 && r[i][j] > r[i-1][j] {
				repoTmp[i][j] = mmax(repoTmp[i][j], repoTmp[i-1][j] + 1)
			}
			if j-1 >= 0 && r[i][j] > r[i][j-1] {
				repoTmp[i][j] = mmax(repoTmp[i][j], repoTmp[i][j-1] + 1)
			}

			if i+1 < len(r) && r[i][j] > r[i+1][j] {
				repoTmp[i][j] = mmax(repoTmp[i][j], repoTmp[i+1][j] + 1)
			}
			if j+1 < len(r[0]) && r[i][j] > r[i][j+1] {
				repoTmp[i][j] = mmax(repoTmp[i][j], repoTmp[i][j+1] + 1)
			}

			res = mmax(res, repoTmp[i][j])
		}
	}
	fmt.Printf("down-right -> up-left: %+v\n", repoTmp)
	fmt.Printf("res is %d\n", res)
	return res
}

func mmax(a, b int) int {
	if a < b { return b }
	return a
}