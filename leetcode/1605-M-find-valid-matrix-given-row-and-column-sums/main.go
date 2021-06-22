package main

func main() {

}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	row, col := len(rowSum), len(colSum)
	res := make([][]int, row)
	for i := range res {
		res[i] = make([]int, col)
	}

	// rowSum: 各行的和
	// colSum: 各列的和
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			res[i][j] = min(rowSum[i], colSum[j])
			rowSum[i] -= res[i][j]
			colSum[j] -= res[i][j]
		}
	}
	return res
}

func min(a, b int) int {
	if a <  b {
		return a
	}
	return b
}