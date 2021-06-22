package main

func main() {

}

type per struct {
	pos int
	neg int
}
func maxProductPath(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	repo := make([][]per, row)
	for k := range repo {
		repo[k] = make([]per, col)
	}
	repo[0][0] = per{grid[0][0], grid[0][0]}
	for j := 1; j < len(grid[0]); j++ {
		repo[0][j] = per{max(repo[0][j-1].pos * grid[0][j], repo[0][j-1].neg * grid[0][j]),
			min(repo[0][j-1].pos * grid[0][j], repo[0][j-1].neg * grid[0][j])}
	}
	for i := 1; i < len(grid); i++ {
		repo[i][0] = per{max(repo[i-1][0].pos * grid[i][0], repo[i-1][0].neg * grid[i][0]),
			min(repo[i-1][0].pos * grid[i][0], repo[i-1][0].neg * grid[i][0])}
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			repo[i][j].pos = findMax(repo[i-1][j], repo[i][j-1], grid[i][j])
			repo[i][j].neg = findMin(repo[i-1][j], repo[i][j-1], grid[i][j])
		}
	}

	val := max(repo[len(grid)-1][len(grid[0])-1].pos, repo[len(grid)-1][len(grid[0])-1].neg)
	if val < 0 {
		return -1
	}
	return val%(1000000000+7)
}

func max(a, b int) int {
	if a <  b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a <  b {
		return a
	}
	return b
}

func findMax(top, left per, val int) int {
	topMax := max(top.pos * val, top.neg * val)
	leftMax := max(left.pos * val, left.neg * val)
	return max(topMax, leftMax)
}

func findMin(top, left per, val int) int {
	topMin := min(top.pos * val, top.neg * val)
	leftMin := min(left.pos * val, left.neg * val)
	return min(topMin, leftMin)
}