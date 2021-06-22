package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	stwo := &TreeNode{2, nil, nil}
	sthree := &TreeNode{3, nil, nil}
	root := &TreeNode{1, stwo, sthree}

	fmt.Println(preorderTraversal(root))

}

func preorderTraversal(root *TreeNode) []int {
	repo := []int{}
	if root == nil {
		return repo
	}
	nodes := []*TreeNode{root}
	for len(nodes) != 0 {
		srcLenInd := len(nodes) - 1
		node := nodes[srcLenInd]

		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}

		repo = append(repo, node.Val)
		nodes = append(nodes[0:srcLenInd], nodes[srcLenInd+1:]...)
	}

	return repo
}
