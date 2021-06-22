package main

import "fmt"

func main() {
	g1 := [][]int{
		{1,1,1,-1,-1},
		{1,1,1,-1,-1},
		{-1,-1,-1,1,1},
		{1,1,1,1,-1},
		{-1,-1,-1,-1,-1},
	}
	//o1 := []int{1,-1,-1,-1,-1}
	g2 := [][]int{
		{-1},
	}
	//o2 := []int{-1}
	g3 := [][]int{
		{1,1,1,1,1,1},
		{-1,-1,-1,-1,-1,-1},
		{1,1,1,1,1,1},
		{-1,-1,-1,-1,-1,-1},
	}
	//o3 := []int{0,1,2,3,4,-1}
	g4 := [][]int{
		{1,1,1},
	}
	fmt.Println(findBall(g1))
	fmt.Println(findBall(g2))
	fmt.Println(findBall(g3))
	fmt.Println(findBall(g4))
}

func findBall(grid [][]int) []int {

	// 1表示左上到右下
	// -1表示右上到左下
	row, col := len(grid), len(grid[0])
	m := corr(col)
	res := make([]int, col)
	for i := range res {
		res[i] = -1
	}
	if col == 1 { return res }

	for i := 0; i < row; i++ {
		for k, v := range m {
			if v == 0 && grid[i][v] == -1 || v == col-1 && grid[i][v] == 1 {
				delete(m, k)
				continue
			}
			if grid[i][v] == 1 && grid[i][v+1] == -1 {
				delete(m, k)
				continue
			}
			if v-1 >= 0 && grid[i][v-1] == 1 && grid[i][v] == -1 {
				delete(m, k)
				continue
			}
			if v+1 < col && grid[i][v] == 1 && grid[i][v+1] == 1 {
				m[k] = v+1
				continue
			}
			if v-1 >= 0 && grid[i][v-1] == -1 && grid[i][v] == -1 {
				m[k] = v-1
			}
		}
	}
	for k, v := range m {
		res[k] = v
	}
	return res
}

func corr(col int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < col; i++ {
		m[i] = i
	}
	return m
}