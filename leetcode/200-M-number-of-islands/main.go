package main

/*DFS
func numIslands(grid [][]byte) int {
	lenGx := len(grid) //行数
	if lenGx == 0 {
		return 0
	}
	lenGy := len(grid[0]) //列数

	visited := make([][]bool, lenGx)
	for i := 0; i < lenGx; i++ {
		visited[i] = make([]bool, lenGy)
	}

	islands := 0
	for i := 0; i < lenGx; i++ {
		for j := 0; j < lenGy; j++ {
			if visited[i][j] == true || grid[i][j] != '1' {
				continue
			}
			islands++
			singtask(i, j, grid, &visited)
		}
	}
	return islands
}

func singtask(cox, coy int, grid [][]byte, visited *[][]bool) {
	if cox < 0 || cox >= len(grid) || coy < 0 || coy >= len(grid[0]) ||
		grid[cox][coy] != '1' || (*visited)[cox][coy] == true {
		return
	}
	(*visited)[cox][coy] = true

	singtask(cox+1, coy, grid, visited)
	singtask(cox, coy+1, grid, visited)
	singtask(cox-1, coy, grid, visited)
	singtask(cox, coy-1, grid, visited)
}
*/

/*BFS*/
type dim struct {
	x int
	y int
}

func numIslands(grid [][]byte) int {
	lenGx := len(grid) // 行数
	if lenGx == 0 {
		return 0
	}
	lenGy := len(grid[0]) // 列数

	visited := make([][]bool, lenGx)
	for i := 0; i < lenGx; i++ {
		visited[i] = make([]bool, lenGy)
	}

	islandNums := 0 // 岛屿数
	loopNodes := make([]dim, 0)
	for i := 0; i < lenGx; i++ {
		for j := 0; j < lenGy; j++ {
			if visited[i][j] == true || grid[i][j] == '0' {
				continue
			}
			loopNodes = append(loopNodes, dim{i, j})
			islandNums++
			for len(loopNodes) != 0 {
				c := loopNodes[0]
				if len(loopNodes) != 1 {
					loopNodes = loopNodes[0:0]
				} else {
					loopNodes = loopNodes[1:]
				}

				singtask(&visited, &loopNodes, c.x+1, c.y, lenGx, lenGy)
				singtask(&visited, &loopNodes, c.x, c.y+1, lenGx, lenGy)
				singtask(&visited, &loopNodes, c.x-1, c.y, lenGx, lenGy)
				singtask(&visited, &loopNodes, c.x, c.y-1, lenGx, lenGy)
			}
		}
	}
	return islandNums
}

func singtask(visited *[][]bool, loops *[]dim, x, y, lenGx, lenGy int) {
	if x < 0 || x > lenGx || y < 0 || y > lenGy || grid[x][y] == '0' || (*visited)[x][y] == true {
		return
	}

	(*visited)[x][y] = true
	(*loops) = append((*loops), dim{x, y})
}

func main() {
	// leetcode testcase
	// [["1", "1", "0", "0", "1"]]
	// [["1"], ["0"], ["1"], ["0"]]
	// [["1","1","1"],["0","1","0"],["1","1","1"]]
	// [["1","1","1", "1", "0"],["1", "1", "0","1","0"],["1","1","0", "0", "0"], ["0", "0", "0", "0", "0"]]
	// [["1","1","0", "0", "0"],["1", "1", "0","0","0"],["0","0","1", "0", "0"], ["0", "0", "0", "1", "1"]]
	// [["1","1","1","1","1","0","1","1","1","1","1","1","1","1","1","0","1","0","1","1"],["0","1","1","1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","1","0"],["1","0","1","1","1","0","0","1","1","0","1","1","1","1","1","1","1","1","1","1"],["1","1","1","1","0","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"],["1","0","0","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"],["1","0","1","1","1","1","1","1","0","1","1","1","0","1","1","1","0","1","1","1"],["0","1","1","1","1","1","1","1","1","1","1","1","0","1","1","0","1","1","1","1"],["1","1","1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","0","1","1"],["1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","1","1","1","1","1"],["1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"],["0","1","1","1","1","1","1","1","0","1","1","1","1","1","1","1","1","1","1","1"],["1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"],["1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"],["1","1","1","1","1","0","1","1","1","1","1","1","1","0","1","1","1","1","1","1"],["1","0","1","1","1","1","1","0","1","1","1","0","1","1","1","1","0","1","1","1"],["1","1","1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","1","1","0"],["1","1","1","1","1","1","1","1","1","1","1","1","1","0","1","1","1","1","0","0"],["1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"],["1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"],["1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1","1"]]
}
