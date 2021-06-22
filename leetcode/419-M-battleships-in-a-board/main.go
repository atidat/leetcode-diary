package main

func main() {
}
func countBattleships(board [][]byte) int {
	var cnt int
	if board[0][0] == 'X' {
		cnt++
	}
	row, col := len(board), len(board[0])
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if board[r][c] == 'X' {
				if r-1 >= 0 && board[r-1][c] == '.' &&
					c-1 >= 0 && board[r][c-1] == '.' {
					cnt++
				}

				if r == 0 {
					if c-1 >= 0 && board[r][c-1] == '.' {
						cnt++
					}
				}
				if c == 0 {
					if r-1 >= 0 && board[r-1][c] == '.' {
						cnt++
					}
				}
			}
		}
	}
	return cnt
}