package main

import (
	"math"
)

func main() {

}

func updateMatrix(mat [][]int) [][]int {
	calc2(calc1(mat))
	return mat
}


// 计算每个点的左和上
func calc1(mat [][]int) [][]int {
	row, col := len(mat), len(mat[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if mat[i][j] != 0 {
				mat[i][j] = math.MaxInt32
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if mat[i][j] == 0 {
				continue
			}

			// 上
			if i-1 >= 0 {
				mat[i][j] = mmin(mat[i][j], mat[i-1][j]+1)
			}
			// 左
			if j-1 >= 0 {
				mat[i][j] = mmin(mat[i][j], mat[i][j-1]+1)
			}
		}
	}
	return mat
}

// 计算每个点的右和下
func calc2(mat [][]int) {
	row, col := len(mat), len(mat[0])
	for i := row-1; i >= 0; i-- {
		for j := col-1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}

			// 下
			if i+1 < row {
				mat[i][j] = mmin(mat[i][j], mat[i+1][j]+1)
			}
			// 右
			if j+1 < col {
				mat[i][j] = mmin(mat[i][j], mat[i][j+1]+1)
			}
		}
	}
}

func mmin(a, b int) int {
	if a > b { return b }
	return a
}