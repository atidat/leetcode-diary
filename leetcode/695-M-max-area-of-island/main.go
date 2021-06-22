package main

func main() {

}


func maxAreaOfIsland(grid [][]int) int {
	res := 0
	row, col := len(grid), len(grid[0])
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			res = mmax(res, dfs(&grid, x, y))
		}
	}
	return res
}

func dfs(g *[][]int, x, y int) int {
	if x < 0 || x >= len(*g) || y < 0 || y >= len((*g)[0]) || (*g)[x][y] == 0 {
		return 0
	}
	(*g)[x][y] = 0
	res := 1
	res += dfs(g, x+1, y)
	res += dfs(g, x-1, y)
	res += dfs(g, x, y+1)
	res += dfs(g, x, y-1)
	return res
}


func mmax(a, b int) int {
	if a < b {
		return b
	}
	return a
}