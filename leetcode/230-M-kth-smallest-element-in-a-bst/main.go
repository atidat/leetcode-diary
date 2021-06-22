package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
func conRecur(root *TreeNode, kth int) bool {
	curNum := 0
	if root != nil {
		curNum++
	}
	if curNum < kth {
		return false
	}
	if root.Left != nil {
		curNum++
	}
	if root.Right != nil {
		curNum++
	}
	if curNum < kth {
		return false
	}
	return true
}
*/
func kthSmallest(root *TreeNode, k int) int {
	repo := make([]int, 0)

	// 中序遍历
	nodes := make([]*TreeNode, 0)
	tmp := root
	for tmp != nil || len(nodes) != 0{
		if tmp != nil {
			nodes = append(nodes, tmp)
			tmp = tmp.Left
		} else {
			tmp = nodes[len(nodes) - 1]
			repo = append(repo, tmp.Val)
			tmp = tmp.Right
		}
	}
	fmt.Println(repo)
	if len(repo) < k {
		return -1
	}
	return repo[k-1]
}

func main() {
}
