package main

func main() {

}
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	res := 1
	exec(root, &res, root.Val)
	return res
}

func exec(node *TreeNode, res *int, m int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		helper(node.Left, m, res)
	}
	if node.Right != nil {
		helper(node.Right, m, res)
	}

}


func helper(node *TreeNode, m int, res *int) {
	if node.Val >= m {
		*res += 1
		exec(node, res, node.Val)
	} else {
		exec(node, res, m)
	}
}