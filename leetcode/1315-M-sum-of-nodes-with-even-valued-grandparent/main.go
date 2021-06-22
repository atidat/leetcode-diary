package main

func main() {

}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 返回祖父节点为偶数的节点值之和
func sumEvenGrandparent(root *TreeNode) int {
	return recur(root)
}

func recur(node *TreeNode) int {
	if node == nil {
		return 0
	}

	calcCur := node.Val % 2 == 0

	var leftSum, rightSum int
	if calcCur && node.Left != nil {
		if  node.Left.Left != nil {
			leftSum += node.Left.Left.Val
		}
		if node.Left.Right != nil {
			leftSum += node.Left.Right.Val
		}
	}
	leftSum += recur(node.Left)
	if calcCur && node.Right != nil {
		if node.Right.Left != nil {
			rightSum += node.Right.Left.Val
		}
		if node.Right.Right != nil {
			rightSum += node.Right.Right.Val
		}
	}
	rightSum += recur(node.Right)
	return leftSum + rightSum
}