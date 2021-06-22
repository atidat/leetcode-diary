package main

import "fmt"

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(t *TreeNode) bool {
	return dfs(t)
}

