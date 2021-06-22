package main

import (
	"fmt"
	"math"
)

func main() {

}


var ori [][]int = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func getMaximumGold(grid [][]int) int {
	vis := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		vis[i] = make([]bool, len(grid[0]))
	}
	res := math.MinInt32
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res = mmax(res, executor(grid, i, j, 0, &vis))
		}
	}
	fmt.Println(res)
	return res
}

func executor(grid [][]int, i, j int, res int, vis *[][]bool) int {
	glen := len(grid)
	g0len := len(grid[0])

	if i >= glen || i < 0 || j >= g0len || j < 0 || grid[i][j] == 0 {
		return res
	}
	if (*vis)[i][j] {
		return res
	}
	(*vis)[i][j] = true

	max := res

	for k := range ori {
		max = mmax(max, executor(grid, i+ori[k][0], j+ori[k][1], res+grid[i][j], vis))
	}
	(*vis)[i][j] = false
	return max
}

func mmax(a, b int) int {
	if a <  b {
		return b
	}
	return a
}