package main

import "fmt"

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
}

func solve(board [][]byte)  {
	outerLen := len(board)
	if outerLen == 0 {
		return
	}
	innerLen := len(board[0])
	if innerLen == 0 {
		return
	}

	fmt.Printf("%c\n", board)
	for i := 0; i < outerLen; i++ {
		for j := 0; j < innerLen; j++ {
			if (i == 0 || i == outerLen-1 ||
				j == 0 || j == innerLen-1) && board[i][j] == 'O' {
				setContinuousO(board, i, j)
			}
		}
	}
	fmt.Printf("%c\n", board)

	for i := 0; i < outerLen; i++ {
		for j := 0; j < innerLen; j++ {
			if board[i][j] == '%' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
	fmt.Printf("%c\n", board)
}

func setContinuousO(board [][]byte, i, j int) {
	board[i][j] = '%'
	if i > 0 && board[i-1][j] == 'O' {
		setContinuousO(board, i-1, j)
	}
	if j > 0 && board[i][j-1] == 'O' {
		setContinuousO(board, i, j-1)
	}
	if i+1 < len(board) && board[i+1][j] == 'O' {
		setContinuousO(board, i+1, j)
	}
	if j+1 < len(board[i]) && board[i][j+1] == 'O' {
		setContinuousO(board, i, j+1)
	}
}