package main

import "go/ast"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	pPos, qPos := -1, -1
	nodes := make([]*TreeNode, 0)

	run := root
	for run != nil {
		if pPos != -1 && qPos != -1 {
			break
		}

		nodes = append(nodes, run)
		if run.Val == p.Val {
			pPos = len(nodes) - 1
		}
		if run.Val == q.Val {
			qPos = len(nodes) - 1
		}

		if run.Left != nil {
			nodes = append(nodes, run.Left)
		} else {
			var tmp *TreeNode
			nodes = append(nodes, tmp)
		}

		if run.Right != nil {
			nodes = append(nodes, run.Right)
		} else {
			var tmp *TreeNode
			nodes = append(nodes, tmp)
		}
	}
}

func main() {
}
