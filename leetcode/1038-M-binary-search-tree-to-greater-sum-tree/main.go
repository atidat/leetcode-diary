package main

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	exec(root, 0)
	return root
}

func exec(node *TreeNode, acc int) int {
	if node == nil {
		return acc
	}

	node.Val += exec(node.Right, acc)
	if node.Left != nil {
		return exec(node.Left, node.Val)
	}
	return node.Val
}